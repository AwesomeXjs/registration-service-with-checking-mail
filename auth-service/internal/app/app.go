package app

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"

	"github.com/AwesomeXjs/libs/pkg/closer"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/interceptors"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/logger"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
}

// New creates a new instance of App and initializes its dependencies.
func New(ctx context.Context) (*App, error) {
	app := &App{}
	err := app.InitDeps(ctx)
	if err != nil {
		logger.Fatal("failed to init deps", zap.Error(err))
	}
	return app, nil
}

// Run starts the gRPC server and ensures proper resource cleanup.
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()
	go func() {
		log.Println("pprof server is running on :6060")
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatalf("failed to start pprof server: %v", err)
		}
	}()

	err := a.RunGRPCServer()
	if err != nil {
		logger.Fatal("failed to run grpc server", zap.Error(err))
	}

	return nil
}

// InitDeps initializes all dependencies required by the App.
func (a *App) InitDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.InitConfig,
		a.initServiceProvider,
		a.initGrpcServer,
	}
	for _, fun := range inits {
		if err := fun(ctx); err != nil {
			logger.Fatal("failed to init deps", zap.Error(err))
		}
	}
	return nil
}

// InitConfig loads environment variables from the specified path.
func (a *App) InitConfig(_ context.Context) error {
	err := godotenv.Load(envPath)
	if err != nil {
		logger.Error("Error loading .env file", zap.String("path", envPath))
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

	a.grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				interceptors.LogInterceptor),
		))
	reflection.Register(a.grpcServer)
	authService.RegisterAuthV1Server(a.grpcServer, a.serviceProvider.GrpcServer(ctx))

	return nil
}

// RunGRPCServer starts the gRPC server and listens on the configured address.
func (a *App) RunGRPCServer() error {
	logger.Info("starting grpc server on " + a.serviceProvider.GRPCConfig().GetAddress())
	list, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().GetAddress())
	if err != nil {
		logger.Fatal("failed to listen grpc", zap.Error(err))
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		logger.Fatal("failed to serve grpc", zap.Error(err))
	}

	return nil
}
