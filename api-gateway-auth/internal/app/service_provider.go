package app

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/client/auth_client"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/configs"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/controller"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/utils/closer"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/utils/header_helper"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/utils/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// serviceProvider manages the application's configuration, clients, and controllers.
type serviceProvider struct {
	// configs
	httpConfig       *configs.HTTPConfig
	authClientConfig *configs.AuthClient

	// clients
	authClient   auth_client.AuthClient
	headerHelper header_helper.IHeaderHelper

	// controllers
	controller *controller.Controller
}

// newServiceProvider creates a new instance of serviceProvider.
func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// HTTPConfig returns the HTTP configuration, initializing it if necessary.
func (s *serviceProvider) HTTPConfig() *configs.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := configs.NewHTTPConfig()
		if err != nil {
			logger.Fatal("failed to get http config", zap.Error(err))
		}
		s.httpConfig = cfg
	}
	return s.httpConfig
}

// AuthClientConfig returns the authentication client configuration, initializing it if necessary.
func (s *serviceProvider) AuthClientConfig() *configs.AuthClient {
	if s.authClientConfig == nil {
		cfg, err := configs.NewAuthClient()
		if err != nil {
			logger.Fatal("failed to get grpc config", zap.Error(err))
		}
		s.authClientConfig = cfg
	}
	return s.authClientConfig
}

// AuthClient returns the authentication client, initializing it if necessary.
func (s *serviceProvider) AuthClient(_ context.Context) auth_client.AuthClient {
	if s.authClient == nil {
		conn, err := grpc.NewClient(s.AuthClientConfig().Address(), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			logger.Fatal(err.Error())
		}
		closer.Add(conn.Close)

		client := auth_v1.NewAuthV1Client(conn)
		s.authClient = auth_client.New(client)
	}
	return s.authClient
}

// HeaderHelper returns the header helper instance, initializing it if necessary.
func (s *serviceProvider) HeaderHelper() header_helper.IHeaderHelper {
	if s.headerHelper == nil {
		s.headerHelper = header_helper.New()
	}
	return s.headerHelper
}

// Controller returns the controller instance, initializing it if necessary.
func (s *serviceProvider) Controller(ctx context.Context) *controller.Controller {
	if s.controller == nil {
		s.controller = controller.New(s.AuthClient(ctx), s.HeaderHelper())
	}
	return s.controller
}
