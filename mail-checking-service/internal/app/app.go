package app

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"sync"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/interceptors"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/pkg/closer"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/pkg/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/pkg/mail_v1"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	EnvPath     = ".env" // EnvPath - contains path to .env file
	serviceName = "mail-checking-service"
)

var (
	// LogLevel defines the logging level, which can be set using the command-line flag "-l".
	LogLevel = flag.String("l", "info", "log level")
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

	wg := &sync.WaitGroup{}

	wg.Add(5)

	go func() {
		defer wg.Done()
		err := a.runPrometheus()
		if err != nil {
			logger.Fatal("failed to run metrics", mark, zap.Error(err))
		}
	}()

	go func() {
		defer wg.Done()
		err := a.RunGRPSServer()
		if err != nil {
			logger.Fatal("failed to run grpc server", mark, zap.Error(err))
		}
	}()

	go func() {
		defer wg.Done()
		cons := a.serviceProvider.KafkaConsumer(ctx, 1)
		cons.Start(ctx)
		closer.Add(cons.Stop)
	}()

	go func() {
		defer wg.Done()
		cons := a.serviceProvider.KafkaConsumer(ctx, 2)
		cons.Start(ctx)
		closer.Add(cons.Stop)
	}()

	go func() {
		defer wg.Done()
		cons := a.serviceProvider.KafkaConsumer(ctx, 3)
		cons.Start(ctx)
		closer.Add(cons.Stop)
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

	err := godotenv.Load(EnvPath)
	if err != nil {
		logger.Error("Error loading .env file", mark, zap.String("path", EnvPath))
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

	a.grpcServer = grpc.NewServer(grpc.UnaryInterceptor(
		grpc_middleware.ChainUnaryServer(
			otgrpc.OpenTracingServerInterceptor(opentracing.GlobalTracer()),
			interceptors.LogInterceptor,
			interceptors.MetricsInterceptor,
		)))
	reflection.Register(a.grpcServer)
	mail_v1.RegisterMailV1Server(a.grpcServer, a.serviceProvider.GrpcServer(ctx))

	return nil
}

// RunGRPSServer starts the gRPC server and listens on the configured address.
func (a *App) RunGRPSServer() error {

	const mark = "App.app.RunGRPSServer"

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
