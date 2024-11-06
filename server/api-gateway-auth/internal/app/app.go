package app

import (
	"context"
	"flag"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/configs"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/closer"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/consts"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

var logLevel = flag.String("l", "info", "log level")

type App struct {
	serviceProvider *serviceProvider
	server          *echo.Echo
}

func New(ctx context.Context) (*App, error) {
	app := &App{}
	err := app.InitDeps(ctx)
	if err != nil {
		logger.Fatal("failed to init deps", zap.Error(err))
	}
	return app, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()
	err := a.runHTTPServer()
	if err != nil {
		logger.Fatal("failed to run http server", zap.Error(err))
	}
	return nil
}

func (a *App) InitDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.InitConfig,
		a.InitEchoServer,
		a.initServiceProvider,
	}
	for _, fun := range inits {
		if err := fun(ctx); err != nil {
			logger.Fatal("failed to init deps", zap.Error(err))
		}
	}
	a.InitRoutes(ctx, a.server)
	return nil
}

func (a *App) InitConfig(_ context.Context) error {
	if err := configs.LoadEnv(consts.EnvPath); err != nil {
		logger.Fatal("failed to load env", zap.Error(err))
	}
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) InitEchoServer(_ context.Context) error {
	flag.Parse()
	logger.Init(logger.GetCore(logger.GetAtomicLevel(logLevel)))

	a.server = echo.New()
	a.server.Use(middleware.Recover())
	a.server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080", "http://localhost:9999"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	return nil
}

func (a *App) runHTTPServer() error {
	logger.Info("server listening at %v", zap.String("start", a.serviceProvider.HTTPConfig().Address()))
	return a.server.Start(a.serviceProvider.HTTPConfig().Address())
}

func (a *App) InitRoutes(ctx context.Context, server *echo.Echo) {
	a.serviceProvider.Controller(ctx).InitRoutes(server)
}
