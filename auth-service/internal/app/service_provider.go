package app

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/clients/db"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/clients/db/pg"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/clients/kafka"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/clients/redis"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/clients/redis/go_redis"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/configs"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/controller"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/repository"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/service"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/utils/auth_helper"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/utils/closer"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/utils/logger"
	"go.uber.org/zap"
)

// serviceProvider struct holds configurations and instances needed to set up and manage services.
type serviceProvider struct {
	// configs
	pgConfig    configs.PGConfig
	grpcConfig  configs.GRPCConfig
	redisConfig configs.RedisConfig
	kafkaConfig configs.KafkaConfig

	// clients
	dbClient      db.Client
	authHelper    auth_helper.AuthHelper
	redisClient   redis.IRedis
	kafkaProducer kafka.IProducer

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

// RedisConfig retrieves the Redis configuration, initializing it if necessary.
func (s *serviceProvider) RedisConfig() configs.RedisConfig {
	if s.redisConfig == nil {
		cfg, err := configs.NewRedisConfig()
		if err != nil {
			logger.Fatal("failed to get redis config", zap.Error(err))
		}
		s.redisConfig = cfg
	}
	return s.redisConfig
}

func (s *serviceProvider) KafkaConfig() configs.KafkaConfig {
	if s.kafkaConfig == nil {
		cfg, err := configs.NewKafkaConfig()
		if err != nil {
			logger.Fatal("failed to get kafka config", zap.Error(err))
		}
		s.kafkaConfig = cfg
	}
	return s.kafkaConfig

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

// RedisClient initializes and returns the Redis client if not already created.
// It also pings Redis to ensure the connection is valid.
func (s *serviceProvider) RedisClient(ctx context.Context) redis.IRedis {
	if s.redisClient == nil {
		redisClient := go_redis.NewGoRedisClient(s.RedisConfig())
		closer.Add(redisClient.Client.Close)

		err := redisClient.Client.Ping(ctx).Err()
		if err != nil {
			logger.Error("Failed to connect to redis", zap.Error(err))
		}

		s.redisClient = redisClient
	}
	return s.redisClient
}

// KafkaProducer returns an instance of the Kafka producer.
// If the Kafka producer is not created yet, it creates a new one using the Kafka configuration addresses.
// In case of an error while creating the producer, the function logs a fatal error and stops execution.
// It also adds the producer's Close method to "closer" to ensure proper cleanup when done.
func (s *serviceProvider) KafkaProducer() kafka.IProducer {
	if s.kafkaProducer == nil {
		producer, err := kafka.NewProducer(s.KafkaConfig().Address())
		if err != nil {
			logger.Fatal("failed to create kafka producer", zap.Error(err))
		}
		closer.Add(producer.Close)
		s.kafkaProducer = producer
	}
	return s.kafkaProducer
}

// AuthHelper initializes and returns the authentication helper if not already created.
func (s *serviceProvider) AuthHelper() auth_helper.AuthHelper {
	if s.authHelper == nil {
		cfg, err := configs.NewAuthConfig()
		if err != nil {
			logger.Fatal("failed to get auth config", zap.Error(err))
		}

		s.authHelper = auth_helper.New(cfg.GetSecretKey(), cfg.GetRefreshTokenDuration(), cfg.GetAccessTokenDuration())
	}
	return s.authHelper
}

// Repository initializes and returns the repository layer for database operations.
func (s *serviceProvider) Repository(ctx context.Context) repository.IRepository {
	if s.repository == nil {
		s.repository = repository.New(s.DBClient(ctx), s.RedisClient(ctx))
	}
	return s.repository
}

// Service initializes and returns the service layer for core business logic.
func (s *serviceProvider) Service(ctx context.Context) service.IService {
	if s.service == nil {
		s.service = service.New(s.Repository(ctx), s.AuthHelper(), s.KafkaProducer())
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
