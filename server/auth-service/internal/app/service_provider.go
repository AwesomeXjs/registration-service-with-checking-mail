package app

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/clients/db"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/clients/db/pg"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/configs"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/controller"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/repository"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/service"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/authHelper"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/closer"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/logger"
	"go.uber.org/zap"
)

// serviceProvider struct holds configurations and instances needed to set up and manage services.
type serviceProvider struct {
	// configs
	pgConfig   configs.PGConfig
	grpcConfig configs.GRPCConfig

	// clients
	dbClient   db.Client
	authHelper authHelper.AuthHelper

	// layers
	controller *controller.Controller
	service    service.IService
	repository repository.IRepository
}

// newServiceProvider creates a new instance of serviceProvider.
func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// PGConfig initializes and returns the PostgreSQL configuration if not already set.
func (s *serviceProvider) PGConfig() configs.PGConfig {
	if s.pgConfig == nil {
		cfg, err := configs.NewPgConfig()
		if err != nil {
			logger.Fatal("failed to get pg config", zap.Error(err))
		}
		s.pgConfig = cfg
	}
	return s.pgConfig
}

// GRPCConfig initializes and returns the gRPC configuration if not already set.
func (s *serviceProvider) GRPCConfig() configs.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := configs.NewGrpcConfig()
		if err != nil {
			logger.Fatal("failed to get grpc config", zap.Error(err))
		}
		s.grpcConfig = cfg
	}
	return s.grpcConfig
}

// DBClient initializes and returns the database client if not already created.
// It also pings the database to ensure the connection is valid.
func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cfg := s.PGConfig()
		dbc, err := pg.New(ctx, cfg.GetDSN())
		if err != nil {
			logger.Fatal("failed to get db client", zap.Error(err))
		}

		err = dbc.DB().Ping(ctx)
		if err != nil {
			logger.Fatal("failed to ping db", zap.Error(err))
		}

		closer.Add(dbc.Close) // Ensures the database client is closed on shutdown
		s.dbClient = dbc
	}
	return s.dbClient
}

func (s *serviceProvider) AuthHelper() authHelper.AuthHelper {
	if s.authHelper == nil {
		cfg, err := configs.NewAuthConfig()
		if err != nil {
			logger.Fatal("failed to get auth config", zap.Error(err))
		}

		s.authHelper = authHelper.New(cfg.GetSecretKey(), cfg.GetRefreshTokenDuration(), cfg.GetAccessTokenDuration())

	}
	return s.authHelper
}

// Repository initializes and returns the repository layer for database operations.
func (s *serviceProvider) Repository(ctx context.Context) repository.IRepository {
	if s.repository == nil {
		s.repository = repository.New(s.DBClient(ctx))
	}
	return s.repository
}

// Service initializes and returns the service layer for core business logic.
func (s *serviceProvider) Service(ctx context.Context) service.IService {
	if s.service == nil {
		s.service = service.New(s.Repository(ctx), s.AuthHelper())
	}
	return s.service
}

// Controller initializes and returns the controller layer to handle business logic requests.
func (s *serviceProvider) Controller(ctx context.Context) *controller.Controller {
	if s.controller == nil {
		s.controller = controller.New(s.Service(ctx))
	}
	return s.controller
}
