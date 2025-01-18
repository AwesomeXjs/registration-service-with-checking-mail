package app

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/metrics"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/clients/kafka"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/clients/redis"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/clients/redis/go_redis"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/grpc_server"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/jwt_manager"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/repository"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/service"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/closer"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/db"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/db/pg"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/db/transaction"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"

	"go.uber.org/zap"
)

// serviceProvider struct holds configurations and instances needed to set up and manage services.
type serviceProvider struct {
	// configs
	pgConfig         db.PGConfig
	grpcConfig       GRPCConfig
	redisConfig      redis.IRedisConfig
	kafkaConfig      kafka.IKafkaConfig
	prometheusConfig metrics.PrometheusConfig

	// clients
	dbClient      db.Client
	txManager     db.TxManager
	authHelper    jwt_manager.AuthHelper
	redisClient   redis.IRedis
	kafkaProducer kafka.IProducer

	// layers
	grpcServer *grpc_server.GrpcServer
	service    *service.Service
	repository *repository.Repository
}

// newServiceProvider creates a new instance of serviceProvider.
func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// PGConfig initializes and returns the PostgresSQL configuration if not already set.
func (s *serviceProvider) PGConfig() db.PGConfig {
	const mark = "App.ServiceProvider.PGConfig"

	if s.pgConfig == nil {
		cfg, err := db.NewPgConfig()
		if err != nil {
			logger.Fatal("failed to get pg config", mark, zap.Error(err))
		}
		s.pgConfig = cfg
	}
	return s.pgConfig
}

// GRPCConfig initializes and returns the gRPC configuration if not already set.
func (s *serviceProvider) GRPCConfig() GRPCConfig {
	const mark = "App.ServiceProvider.GRPCConfig"

	if s.grpcConfig == nil {
		cfg, err := NewGrpcConfig()
		if err != nil {
			logger.Fatal("failed to get grpc config", mark, zap.Error(err))
		}
		s.grpcConfig = cfg
	}
	return s.grpcConfig
}

// RedisConfig retrieves the Redis configuration, initializing it if necessary.
func (s *serviceProvider) RedisConfig() redis.IRedisConfig {
	const mark = "App.ServiceProvider.RedisConfig"

	if s.redisConfig == nil {
		cfg, err := redis.NewRedisConfig()
		if err != nil {
			logger.Fatal("failed to get redis config", mark, zap.Error(err))
		}
		s.redisConfig = cfg
	}
	return s.redisConfig
}

func (s *serviceProvider) KafkaConfig() kafka.IKafkaConfig {
	const mark = "App.ServiceProvider.KafkaConfig"

	if s.kafkaConfig == nil {
		cfg, err := kafka.NewKafkaConfig()
		if err != nil {
			logger.Fatal("failed to get kafka config", mark, zap.Error(err))
		}
		s.kafkaConfig = cfg
	}
	return s.kafkaConfig

}

func (s *serviceProvider) PrometheusConfig() metrics.PrometheusConfig {

	const mark = "App.ServiceProvider.PrometheusConfig"

	if s.prometheusConfig == nil {
		cfg, err := metrics.NewPrometheusConfig()
		if err != nil {
			logger.Fatal("failed to get metrics config", mark, zap.Error(err))
		}
		s.prometheusConfig = cfg
	}

	return s.prometheusConfig
}

// DBClient initializes and returns the database client if not already created.
// It also pings the database to ensure the connection is valid.
func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	const mark = "App.ServiceProvider.DBClient"

	if s.dbClient == nil {
		cfg := s.PGConfig()
		dbc, err := pg.New(ctx, cfg.GetDSN())
		if err != nil {
			logger.Fatal("failed to get db client", mark, zap.Error(err))
		}

		err = dbc.DB().Ping(ctx)
		if err != nil {
			logger.Fatal("failed to ping db", mark, zap.Error(err))
		}

		closer.Add(dbc.Close) // Ensures the database client is closed on shutdown
		s.dbClient = dbc
	}
	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}
	return s.txManager
}

// RedisClient initializes and returns the Redis client if not already created.
// It also pings Redis to ensure the connection is valid.
func (s *serviceProvider) RedisClient(ctx context.Context) redis.IRedis {
	const mark = "App.ServiceProvider.RedisClient"

	if s.redisClient == nil {
		redisClient := go_redis.NewGoRedisClient(s.RedisConfig())
		closer.Add(redisClient.Client.Close)

		err := redisClient.Client.Ping(ctx).Err()
		if err != nil {
			logger.Error("Failed to connect to redis", mark, zap.Error(err))
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
	const mark = "App.ServiceProvider.KafkaProducer"

	if s.kafkaProducer == nil {
		producer, err := kafka.NewProducer(s.KafkaConfig().Address())
		if err != nil {
			logger.Fatal("failed to create kafka producer", mark, zap.Error(err))
		}
		closer.Add(producer.Close)
		s.kafkaProducer = producer
	}
	return s.kafkaProducer
}

// AuthHelper initializes and returns the authentication helper if not already created.
func (s *serviceProvider) AuthHelper() jwt_manager.AuthHelper {
	const mark = "App.ServiceProvider.AuthHelper"

	if s.authHelper == nil {
		cfg, err := jwt_manager.NewAuthConfig()
		if err != nil {
			logger.Fatal("failed to get auth config", mark, zap.Error(err))
		}

		s.authHelper = jwt_manager.New(cfg.GetSecretKey(), cfg.GetRefreshTokenDuration(), cfg.GetAccessTokenDuration())
	}
	return s.authHelper
}

// Repository initializes and returns the repository layer for database operations.
func (s *serviceProvider) Repository(ctx context.Context) *repository.Repository {
	if s.repository == nil {
		s.repository = repository.New(s.DBClient(ctx), s.RedisClient(ctx))
	}
	return s.repository
}

// Service initializes and returns the service layer for core business logic.
func (s *serviceProvider) Service(ctx context.Context) *service.Service {
	if s.service == nil {
		s.service = service.New(s.Repository(ctx),
			s.AuthHelper(),
			s.TxManager(ctx),
			s.KafkaProducer(),
			s.DBClient(ctx))
	}
	return s.service
}

// GrpcServer initializes and returns the controller layer to handle business logic requests.
func (s *serviceProvider) GrpcServer(ctx context.Context) *grpc_server.GrpcServer {
	if s.grpcServer == nil {
		s.grpcServer = grpc_server.New(s.Service(ctx))
	}
	return s.grpcServer
}
