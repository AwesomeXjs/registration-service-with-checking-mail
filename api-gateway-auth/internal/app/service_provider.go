package app

import (
	"context"

	"github.com/AwesomeXjs/libs/pkg/closer"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/client/grpc_auth_client"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/client/mail_client"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/controller"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/headers_manager"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/pkg/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/pkg/mail_v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// serviceProvider manages the application's configuration, clients, and controllers.
type serviceProvider struct {
	// configs
	httpConfig       IHTTPConfig
	authClientConfig grpc_auth_client.IAuthClientConfig
	mailClientConfig mail_client.IMailClientConfig

	// clients
	authClient   grpc_auth_client.AuthClient
	mailClient   mail_client.MailClient
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

	const mark = "App.ServiceProvider.HTTPConfig"

	if s.httpConfig == nil {
		cfg, err := NewHTTPConfig()
		if err != nil {
			logger.Fatal("failed to get http config", mark, zap.Error(err))
		}
		s.httpConfig = cfg
	}
	return s.httpConfig
}

// AuthClientConfig returns the authentication client configuration, initializing it if necessary.
func (s *serviceProvider) AuthClientConfig() grpc_auth_client.IAuthClientConfig {

	const mark = "App.ServiceProvider.AuthClientConfig"

	if s.authClientConfig == nil {
		cfg, err := grpc_auth_client.NewAuthClient()
		if err != nil {
			logger.Fatal("failed to get grpc config", mark, zap.Error(err))
		}
		s.authClientConfig = cfg
	}
	return s.authClientConfig
}

func (s *serviceProvider) MailClientConfig() mail_client.IMailClientConfig {

	const mark = "App.ServiceProvider.MailClientConfig"

	if s.mailClientConfig == nil {
		cfg, err := mail_client.NewMailClient()
		if err != nil {
			logger.Fatal("failed to get grpc config", mark, zap.Error(err))
		}
		s.mailClientConfig = cfg
	}
	return s.mailClientConfig
}

// GrpcAuthClient returns the authentication client, initializing it if necessary.
func (s *serviceProvider) GrpcAuthClient(_ context.Context) grpc_auth_client.AuthClient {

	const mark = "App.ServiceProvider.GrpcAuthClient"

	if s.authClient == nil {
		conn, err := grpc.NewClient(s.AuthClientConfig().Address(), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			logger.Fatal(err.Error(), mark)
		}
		closer.Add(conn.Close)

		client := auth_v1.NewAuthV1Client(conn)
		s.authClient = grpc_auth_client.New(client)
	}
	return s.authClient
}

func (s *serviceProvider) MailClient() mail_client.MailClient {

	const mark = "App.ServiceProvider.MailClient"

	if s.mailClient == nil {
		conn, err := grpc.NewClient(s.MailClientConfig().Address(), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			logger.Fatal(err.Error(), mark)
		}
		closer.Add(conn.Close)

		client := mail_v1.NewMailV1Client(conn)
		s.mailClient = mail_client.NewGRPCMailClient(client)
	}
	return s.mailClient
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
		s.controller = controller.New(s.GrpcAuthClient(ctx), s.MailClient(), s.HeaderHelper())
	}
	return s.controller
}
