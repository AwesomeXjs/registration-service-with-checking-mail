package app

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof" // pprof package
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/joho/godotenv"
	"github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/interceptors"
	ratelimiter "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/rate_limiter"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/closer"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
)

const (
	serviceName = "auth-service"
)

var (
	// LogLevel defines the logging level, which can be set using the command-line flag "-l".
	LogLevel = flag.String("l", "info", "log level")
	envPath  = ".env" // EnvPath - contains path to .env file
)

// App struct encapsulates the dependencies and the gRPC server instance.
type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
	prometheus      *http.Server
}

// New creates a new instance of App and initializes its dependencies.
func New(ctx context.Context) (*App, error) {
	const mark = "App.app.New"

	app := &App{}
	err := app.InitDeps(ctx)
	if err != nil {
		logger.Fatal("failed to init deps", mark, zap.Error(err))
	}
	return app, nil
}

// Run starts the gRPC server and ensures proper resource cleanup.
func (a *App) Run(ctx context.Context) error {
	const mark = "App.app.Run"

	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	// outbox scheduler
	a.serviceProvider.Service(ctx).Event.Start(ctx, GetSchedulerPeriod())

	wg := &sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Wait()
		logger.Info("starting pprof server on :6060", mark)
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatalf("failed to start pprof server: %v", err)
		}
	}()

	go func() {
		defer wg.Done()
		err := a.runPrometheus()
		if err != nil {
			logger.Fatal("failed to run metrics", mark, zap.Error(err))
		}
	}()
	go func() {
		defer wg.Done()
		err := a.RunGRPCServer()
		if err != nil {
			logger.Fatal("failed to run grpc server", mark, zap.Error(err))
		}
	}()

	wg.Wait()
	return nil
}

// InitDeps initializes all dependencies required by the App.
func (a *App) InitDeps(ctx context.Context) error {
	const mark = "App.app.InitDeps"

	inits := []func(context.Context) error{
		a.InitConfig,
		a.initServiceProvider,
		a.initGrpcServer,
		a.initPrometheus,
		a.initMetrics,
	}
	for _, fun := range inits {
		if err := fun(ctx); err != nil {
			logger.Fatal("failed to init deps", mark, zap.Error(err))
		}
	}
	return nil
}

// InitConfig loads environment variables from the specified path.
func (a *App) InitConfig(_ context.Context) error {
	const mark = "App.app.InitConfig"

	err := godotenv.Load(envPath)
	if err != nil {
		logger.Error("Error loading .env file", mark, zap.String("path", envPath))
		return fmt.Errorf("error loading .env file: %v", err)
	}
	return nil
}

// initServiceProvider initializes the service provider.
func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

// initGrpcServer initializes the gRPC server and configures logging.
func (a *App) initGrpcServer(ctx context.Context) error {
	flag.Parse()
	logger.Init(logger.GetCore(logger.GetAtomicLevel(LogLevel)))
	a.InitTracing(serviceName)

	rateLimiter := ratelimiter.NewTokenBucketLimiter(ctx, 1000, time.Second)

	a.grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				otgrpc.OpenTracingServerInterceptor(opentracing.GlobalTracer()),
				interceptors.NewRateLimitInterceptor(rateLimiter).Unary,
				interceptors.NewCircuitBreaker(GetCircuitBreakerConfig()).Unary,
				interceptors.LogInterceptor,
				interceptors.MetricsInterceptor,
			),
		))
	reflection.Register(a.grpcServer)
	authService.RegisterAuthV1Server(a.grpcServer, a.serviceProvider.GrpcServer(ctx))

	return nil
}

// RunGRPCServer starts the gRPC server and listens on the configured address.
func (a *App) RunGRPCServer() error {
	const mark = "App.app.RunGRPCServer"

	logger.Info("starting grpc server on "+a.serviceProvider.GRPCConfig().GetAddress(), mark)
	list, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().GetAddress())
	if err != nil {
		logger.Fatal("failed to listen grpc", mark, zap.Error(err))
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		logger.Fatal("failed to serve grpc", mark, zap.Error(err))
	}

	return nil
}

// GetSchedulerPeriod returns the period for the outbox scheduler in milliseconds.
func GetSchedulerPeriod() time.Duration {
	const mark = "App.app.GetSchedulerPeriod"

	period, err := strconv.Atoi(os.Getenv("OUTBOX_SCHEDULER_PERIOD"))
	if err != nil {
		logger.Error("failed to get outbox scheduler period", mark, zap.Error(err))
		log.Fatalf("failed to get outbox scheduler period: %v", err)
	}
	return time.Duration(period) * time.Millisecond
}

// GetCircuitBreakerConfig initializes and returns a configured Circuit Breaker for the auth-service.
// The Circuit Breaker trips when the failure ratio exceeds 60%, with a timeout of 5 seconds for resetting.
func GetCircuitBreakerConfig() *gobreaker.CircuitBreaker {
	const mark = "App.app.GetCircuitBreakerConfig"

	cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        serviceName,
		MaxRequests: 3,
		Timeout:     5 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return failureRatio >= 0.6
		},
		OnStateChange: func(name string, from, to gobreaker.State) {
			logger.Warn("circuit breaker state changed", mark, zap.String("name", name), zap.String("from", from.String()), zap.String("to", to.String()))
		},
	})

	return cb
}
