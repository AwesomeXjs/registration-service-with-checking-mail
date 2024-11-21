package app

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/client/kafka"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/client/redis"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/client/redis/go_redis"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/configs"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/controller"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/utils/closer"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/utils/logger"
	"go.uber.org/zap"
)

const (
	topicName = "registration"
)

type serviceProvider struct {
	grpcConfig  configs.GRPCConfigs
	kafkaConfig configs.KafkaConfig
	redisConfig configs.RedisConfig
	emailConfig configs.IMailConfig

	kafkaConsumer *kafka.Consumer
	redisClient   redis.IRedis

	controller   *controller.Controller
	kafkaHandler *kafka.KafkaHandler
}

// newServiceProvider creates a new instance of serviceProvider.
func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// GRPCConfig initializes and returns the gRPC configuration if not already set.
func (s *serviceProvider) GRPCConfig() configs.GRPCConfigs {
	if s.grpcConfig == nil {
		cfg, err := configs.NewGRPCConfig()
		if err != nil {
			logger.Fatal("failed to get grpc config", zap.Error(err))
		}
		s.grpcConfig = cfg
	}
	return s.grpcConfig
}

func (s *serviceProvider) EmailConfig() configs.IMailConfig {
	if s.emailConfig == nil {
		cfg, err := configs.NewMailConfig()
		if err != nil {
			logger.Warn("failed to get email config", zap.Error(err))
		}
		s.emailConfig = cfg
	}
	return s.emailConfig
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

// Controller initializes and returns the controller layer to handle business logic requests.
func (s *serviceProvider) Controller(ctx context.Context) *controller.Controller {
	if s.controller == nil {
		s.controller = controller.New()
	}
	return s.controller
}

func (s *serviceProvider) KafkaHandler(ctx context.Context) *kafka.KafkaHandler {
	if s.kafkaHandler == nil {
		handler := kafka.NewKafkaHandler(s.RedisClient(ctx), s.EmailConfig())
		s.kafkaHandler = handler
	}
	return s.kafkaHandler
}

func (s *serviceProvider) KafkaConsumer(ctx context.Context, number int) *kafka.Consumer {
	if s.kafkaConsumer == nil {
		consumer, err := kafka.NewConsumer(s.KafkaHandler(ctx),
			s.KafkaConfig().Address(), topicName, "mail", number)
		if err != nil {
			logger.Fatal("failed to get kafka consumer", zap.Error(err))
		}
		s.kafkaConsumer = consumer

	}
	return s.kafkaConsumer
}
