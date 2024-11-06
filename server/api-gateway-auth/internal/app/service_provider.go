package app

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/client/auth_client"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/configs"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/controller"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/closer"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/header_helper"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type serviceProvider struct {
	// configs
	httpConfig       *configs.HttpConfig
	authClientConfig *configs.AuthClient

	// clients
	authClient   auth_client.AuthClient
	headerHelper header_helper.IHeaderHelper

	// controllers
	controller *controller.Controller
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) HTTPConfig() *configs.HttpConfig {
	if s.httpConfig == nil {
		cfg, err := configs.NewHTTPConfig()
		if err != nil {
			logger.Fatal("failed to get http config", zap.Error(err))
		}
		s.httpConfig = cfg
	}

	return s.httpConfig
}

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

func (s *serviceProvider) HeaderHelper() header_helper.IHeaderHelper {
	if s.headerHelper == nil {
		s.headerHelper = header_helper.New()
	}
	return s.headerHelper
}

func (s *serviceProvider) Controller(ctx context.Context) *controller.Controller {
	if s.controller == nil {
		s.controller = controller.New(s.AuthClient(ctx), s.HeaderHelper())
	}
	return s.controller
}
