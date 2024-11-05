package configs

import (
	"fmt"
	"net"
	"os"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/logger"
	"go.uber.org/zap"
)

var (
	AuthHost = "GRPC_AUTH_HOST"
	AuthPort = "GRPC_AUTH_PORT"
)

type IAuthClient interface {
	Address()
}

type AuthClient struct {
	Host string
	Port string
}

func NewAuthClient() (*AuthClient, error) {
	host := os.Getenv(AuthHost)
	if len(host) == 0 {
		logger.Error("failed to get auth host", zap.String("auth host", AuthHost))
		return nil, fmt.Errorf("env %s is empty", AuthHost)
	}

	port := os.Getenv(AuthPort)
	if len(port) == 0 {
		logger.Error("failed to get auth port", zap.String("auth port", AuthPort))
		return nil, fmt.Errorf("env %s is empty", AuthPort)
	}

	return &AuthClient{
		Host: host,
		Port: port,
	}, nil
}

func (a *AuthClient) Address() string {
	return net.JoinHostPort(a.Host, a.Port)
}
