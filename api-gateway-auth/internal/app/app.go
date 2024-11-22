package app

import (
	"context"
	"flag"
	"fmt"

	"github.com/AwesomeXjs/libs/pkg/closer"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/middlewares"
	"github.com/asaskevich/govalidator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

const (
	// EnvPath is the path to the .env file that contains environment variables.
	EnvPath = ".env"
)

// logLevel is a command-line flag for specifying the log level.
var logLevel = flag.String("l", "info", "log level")

// App represents the application with dependencies and server.
type App struct {
	serviceProvider *serviceProvider
	server          *echo.Echo
}

// New creates and initializes the App with dependencies.
func New(ctx context.Context) (*App, error) {
	app := &App{}
	err := app.InitDeps(ctx)
	if err != nil {
		// Fatal log in case of failure during dependency initialization
		logger.Fatal("failed to init deps", zap.Error(err))
	}
	return app, nil
}

// Run starts the HTTP server and handles cleanup on shutdown.
func (a *App) Run() error {
	defer func() {
		closer.CloseAll() // Close all services/resources
		closer.Wait()     // Wait for all services to close
	}()
	err := a.runHTTPServer() // Run the HTTP server
	if err != nil {
		logger.Fatal("failed to run http server", zap.Error(err))
	}
	return nil
}

// InitDeps initializes the application's dependencies.
func (a *App) InitDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.InitConfig,          // Initialize config
		a.InitEchoServer,      // Initialize Echo server
		a.initServiceProvider, // Initialize service provider
	}
	for _, fun := range inits {
		if err := fun(ctx); err != nil {
			// Log fatal error if any dependency initialization fails
			logger.Fatal("failed to init deps", zap.Error(err))
		}
	}
	a.InitRoutes(ctx, a.server) // Initialize the routes
	return nil
}

// InitConfig loads environment variables for the application.
func (a *App) InitConfig(_ context.Context) error {
	err := godotenv.Load(EnvPath)
	if err != nil {
		logger.Error("Error loading .env file", zap.String("path", EnvPath))
		return fmt.Errorf("error loading .env file: %v", err)
	}
	return err
}

// initServiceProvider initializes the service provider.
func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider() // Create a new service provider
	return nil
}

// InitEchoServer sets up the Echo server and its middleware.
func (a *App) InitEchoServer(_ context.Context) error {
	flag.Parse()                                                 // Parse command-line flags
	logger.Init(logger.GetCore(logger.GetAtomicLevel(logLevel))) // Initialize logger with the specified log level

	// Custom validator for role enumeration
	govalidator.TagMap["role_enum"] = govalidator.Validator(func(str string) bool {
		validRoles := []string{"admin", "user"} // Valid roles for user
		return govalidator.IsIn(str, validRoles...)
	})

	a.server = echo.New()              // Create a new Echo server
	a.server.Use(middleware.Recover()) // Middleware for recovering from panics
	a.server.Use(middlewares.Logger)   // Custom logging middleware
	a.server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},                                           // Allow CORS from this origin
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE}, // Allowed HTTP methods
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAccessControlAllowCredentials,
			echo.HeaderAuthorization,
			echo.HeaderAccessControlRequestHeaders,
		}, // Allowed headers for CORS
	}))
	return nil
}

// runHTTPServer starts the Echo server and listens for requests.
func (a *App) runHTTPServer() error {
	logger.Info("server listening at %v", zap.String("start", a.serviceProvider.HTTPConfig().Address())) // Log the server address
	return a.server.Start(a.serviceProvider.HTTPConfig().Address())                                      // Start the server at the configured address
}

// InitRoutes sets up the application routes.
func (a *App) InitRoutes(ctx context.Context, server *echo.Echo) {
	a.serviceProvider.Controller(ctx).InitRoutes(server) // Initialize routes using the controller
}
