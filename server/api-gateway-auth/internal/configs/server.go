package configs

import (
	"fmt"
	"net"
	"os"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/logger"
	"go.uber.org/zap"
)

var (
	HttpHost = "HTTP_HOST"
	HttpPort = "HTTP_PORT"
)

type HTTPConfig interface {
	Address() string
}

type HttpConfig struct {
	host string
	port string
}

func NewHTTPConfig() (*HttpConfig, error) {
	host := os.Getenv(HttpHost)
	if len(host) == 0 {
		logger.Error("failed to get http host", zap.String("http host", HttpHost))
		return nil, fmt.Errorf("env %s is empty", HttpHost)
	}

	port := os.Getenv(HttpPort)
	if len(port) == 0 {
		logger.Error("failed to get http port", zap.String("http port", "HTTP_PORT"))
		return nil, fmt.Errorf("env %s is empty", "HTTP_PORT")
	}
	return &HttpConfig{host: host, port: port}, nil
}

func (h *HttpConfig) Address() string {
	return net.JoinHostPort(h.host, h.port)
}
