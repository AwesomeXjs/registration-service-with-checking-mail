package app

import (
	"context"
	"flag"
	"net"
	"sync"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/configs"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/utils/closer"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/utils/consts"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/utils/interceptors"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/utils/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/pkg/mail_v1"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	// LogLevel defines the logging level, which can be set using the command-line flag "-l".
	LogLevel = flag.String("l", "info", "log level")
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

	wg := &sync.WaitGroup{}

	wg.Add(4)
	go func() {
		defer wg.Done()
		err := a.RunGRPSServer()
		if err != nil {
			logger.Fatal("failed to run grpc server", zap.Error(err))
		}
	}()

	go func() {
		defer wg.Done()
		cons := a.serviceProvider.KafkaConsumer(1)
		cons.Start()
		closer.Add(cons.Stop)
	}()

	go func() {
		defer wg.Done()
		cons := a.serviceProvider.KafkaConsumer(2)
		cons.Start()
		closer.Add(cons.Stop)
	}()

	go func() {
		defer wg.Done()
		cons := a.serviceProvider.KafkaConsumer(3)
		cons.Start()
		closer.Add(cons.Stop)
	}()

	wg.Wait()

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
	if err := configs.LoadEnv(consts.EnvPath); err != nil {
		logger.Fatal("failed to load env", zap.Error(err))
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

	a.grpcServer = grpc.NewServer(grpc.UnaryInterceptor(
		grpc_middleware.ChainUnaryServer(
			interceptors.LogInterceptor),
	))
	reflection.Register(a.grpcServer)
	mail_v1.RegisterMailV1Server(a.grpcServer, a.serviceProvider.Controller(ctx))

	return nil
}

// RunGRPSServer starts the gRPC server and listens on the configured address.
func (a *App) RunGRPSServer() error {
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
