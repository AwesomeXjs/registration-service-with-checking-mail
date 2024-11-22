package app

import (
	"context"

	"github.com/AwesomeXjs/libs/pkg/closer"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/client/grpc_auth_client"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/controller"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/headers_manager"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// serviceProvider manages the application's configuration, clients, and controllers.
type serviceProvider struct {
	// configs
	httpConfig       IHTTPConfig
	authClientConfig grpc_auth_client.IAuthClientConfig

	// clients
	authClient   grpc_auth_client.AuthClient
	headerHelper headers_manager.IHeaderHelper

	// controllers
	controller *controller.Controller
}

// newServiceProvider creates a new instance of serviceProvider.
func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// HTTPConfig returns the HTTP configuration, initializing it if necessary.
func (s *serviceProvider) HTTPConfig() IHTTPConfig {
	if s.httpConfig == nil {
		cfg, err := NewHTTPConfig()
		if err != nil {
			logger.Fatal("failed to get http config", zap.Error(err))
		}
		s.httpConfig = cfg
	}
	return s.httpConfig
}

// AuthClientConfig returns the authentication client configuration, initializing it if necessary.
func (s *serviceProvider) AuthClientConfig() grpc_auth_client.IAuthClientConfig {
	if s.authClientConfig == nil {
		cfg, err := grpc_auth_client.NewAuthClient()
		if err != nil {
			logger.Fatal("failed to get grpc config", zap.Error(err))
		}
		s.authClientConfig = cfg
	}
	return s.authClientConfig
}

// GrpcAuthClient returns the authentication client, initializing it if necessary.
func (s *serviceProvider) GrpcAuthClient(_ context.Context) grpc_auth_client.AuthClient {
	if s.authClient == nil {
		conn, err := grpc.NewClient(s.AuthClientConfig().Address(), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			logger.Fatal(err.Error())
		}
		closer.Add(conn.Close)

		client := auth_v1.NewAuthV1Client(conn)
		s.authClient = grpc_auth_client.New(client)
	}
	return s.authClient
}

// HeaderHelper returns the header helper instance, initializing it if necessary.
func (s *serviceProvider) HeaderHelper() headers_manager.IHeaderHelper {
	if s.headerHelper == nil {
		s.headerHelper = headers_manager.New()
	}
	return s.headerHelper
}

// Controller returns the controller instance, initializing it if necessary.
func (s *serviceProvider) Controller(ctx context.Context) *controller.Controller {
	if s.controller == nil {
		s.controller = controller.New(s.GrpcAuthClient(ctx), s.HeaderHelper())
	}
	return s.controller
}
