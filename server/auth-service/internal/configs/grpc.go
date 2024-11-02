package configs

import (
	"fmt"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/consts"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/logger"
	"go.uber.org/zap"
	"net"
	"os"
)

type GRPCConfig interface {
	GetAddress() string
}

type grpcConfig struct {
	host string
	port string
}

func NewGrpcConfig() (GRPCConfig, error) {
	host := os.Getenv(consts.GrpcHost)
	if len(host) == 0 {
		logger.Error("failed to get grpc host", zap.String("grpc host", consts.GrpcHost))
		return nil, fmt.Errorf("env %s is empty", consts.GrpcHost)
	}
	port := os.Getenv(consts.GrpcPort)
	if len(port) == 0 {
		logger.Error("failed to get grpc port", zap.String("grpc port", consts.GrpcPort))
		return nil, fmt.Errorf("env %s is empty", consts.GrpcPort)
	}
	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}

func (g *grpcConfig) GetAddress() string {
	return net.JoinHostPort(g.host, g.port)
}
