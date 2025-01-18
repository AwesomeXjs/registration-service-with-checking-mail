package app

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/metrics"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/client/grpc_auth_client"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/client/kafka"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/client/mail"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/client/redis"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/client/redis/go_redis"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/grpc_server"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/pkg/closer"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/pkg/logger"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	topicName = "registration"
)

type serviceProvider struct {
	grpcConfig       IGRPCConfigs
	kafkaConfig      kafka.IKafkaConfig
	redisConfig      redis.IRedisConfig
	emailConfig      mail.IMailConfig
	authClientConfig grpc_auth_client.IAuthClientConfig
	prometheusConfig metrics.PrometheusConfig

	kafkaConsumer *kafka.Consumer
	redisClient   redis.IRedis
	mailClient    mail.IMailClient
	authClient    grpc_auth_client.IAuthClient

	grpcServer   *grpc_server.GrpcServer
	kafkaHandler kafka.IKafkaHandler
}

// newServiceProvider creates a new instance of serviceProvider.
func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// GRPCConfig initializes and returns the gRPC configuration if not already set.
func (s *serviceProvider) GRPCConfig() IGRPCConfigs {

	const mark = "App.ServiceProvider.GRPCConfig"

	if s.grpcConfig == nil {
		cfg, err := NewGRPCConfig()
		if err != nil {
			logger.Fatal("failed to get grpc config", mark, zap.Error(err))
		}
		s.grpcConfig = cfg
	}
	return s.grpcConfig
}

func (s *serviceProvider) MailConfig() mail.IMailConfig {

	const mark = "App.ServiceProvider.MailConfig"

	if s.emailConfig == nil {
		cfg, err := mail.NewMailConfig()
		if err != nil {
			logger.Warn("failed to get email config", mark, zap.Error(err))
		}
		s.emailConfig = cfg
	}
	return s.emailConfig
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

func (s *serviceProvider) AuthClientConfig() grpc_auth_client.IAuthClientConfig {

	const mark = "App.ServiceProvider.AuthClientConfig"

	if s.authClientConfig == nil {
		cfg, err := grpc_auth_client.NewAuthClientConfig()
		if err != nil {
			logger.Fatal("failed to get grpc config", mark, zap.Error(err))
		}
		s.authClientConfig = cfg
	}
	return s.authClientConfig
}

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

func (s *serviceProvider) MailClient() mail.IMailClient {
	if s.mailClient == nil {
		mailClient := mail.NewMailClient(s.MailConfig())

		s.mailClient = mailClient
	}
	return s.mailClient
}

func (s *serviceProvider) AuthClient() grpc_auth_client.IAuthClient {

	const mark = "App.ServiceProvider.AuthClient"

	if s.authClient == nil {
		conn, err := grpc.NewClient(s.AuthClientConfig().Address(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())))
		if err != nil {
			logger.Fatal(err.Error(), mark)
		}
		closer.Add(conn.Close)

		client := auth_v1.NewAuthV1Client(conn)
		s.authClient = grpc_auth_client.NewAuthClient(client)
	}
	return s.authClient
}

func (s *serviceProvider) GrpcServer(ctx context.Context) *grpc_server.GrpcServer {
	if s.grpcServer == nil {
		s.grpcServer = grpc_server.New(s.RedisClient(ctx), s.AuthClient())
	}
	return s.grpcServer
}

func (s *serviceProvider) KafkaHandler(ctx context.Context) kafka.IKafkaHandler {
	if s.kafkaHandler == nil {
		handler := kafka.NewKafkaHandler(s.RedisClient(ctx), s.MailClient())
		s.kafkaHandler = handler
	}
	return s.kafkaHandler
}

func (s *serviceProvider) KafkaConsumer(ctx context.Context, number int) *kafka.Consumer {

	const mark = "App.ServiceProvider.KafkaConsumer"

	if s.kafkaConsumer == nil {
		consumer, err := kafka.NewConsumer(s.KafkaHandler(ctx),
			s.KafkaConfig().Address(), topicName, "mail", number)
		if err != nil {
			logger.Fatal("failed to get kafka consumer", mark, zap.Error(err))
		}
		s.kafkaConsumer = consumer

	}
	return s.kafkaConsumer
}
