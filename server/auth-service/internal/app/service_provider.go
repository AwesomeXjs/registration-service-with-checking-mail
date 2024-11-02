package app

import (
	"context"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/clients/db"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/clients/db/pg"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/configs"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/controller"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/repository"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/service"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/closer"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/logger"
	"go.uber.org/zap"
)

type serviceProvider struct {
	pgConfig   configs.PGConfig
	grpcConfig configs.GRPCConfig

	dbClient db.Client

	controller *controller.Controller
	service    service.IService
	repository repository.IRepository
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

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

		closer.Add(dbc.Close)
		s.dbClient = dbc
	}
	return s.dbClient
}

func (s *serviceProvider) Repository(ctx context.Context) repository.IRepository {
	if s.repository == nil {
		s.repository = repository.New(s.DBClient(ctx))
	}
	return s.repository
}

func (s *serviceProvider) Service(ctx context.Context) service.IService {
	if s.service == nil {
		s.service = service.New(s.Repository(ctx))
	}
	return s.service
}

func (s *serviceProvider) Controller(ctx context.Context) *controller.Controller {
	if s.controller == nil {
		s.controller = controller.New(s.Service(ctx))
	}
	return s.controller
}
