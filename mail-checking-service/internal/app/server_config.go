package app

import (
	"fmt"
	"net"
	"os"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/logger"
	"go.uber.org/zap"
)

const (
	// GrpcHost is the environment variable key for the gRPC server hostname.
	// It specifies where the gRPC server can be reached, set in the .env file.
	GrpcHost = "GRPC_HOST"

	// GrpcPort is the environment variable key for the gRPC server port number.
	// It indicates the port on which the gRPC server listens for connections, also set in the .env file.
	GrpcPort = "GRPC_PORT"
)

// GRPCConfig defines an interface for obtaining the gRPC server address.
type GRPCConfigs interface {
	GetAddress() string
}

// grpcConfig implements the GRPCConfig interface, storing the host and port for the gRPC server.
type grpcConfig struct {
	host string
	port string
}

func (g *grpcConfig) GetAddress() string {
	return net.JoinHostPort(g.host, g.port)
}

// NewGRPCConfig creates a new GRPCConfig instance by reading the host and port from environment variables.
// It returns an error if either value is not set.
func NewGRPCConfig() (GRPCConfigs, error) {
	host := os.Getenv(GrpcHost)
	if len(host) == 0 {
		logger.Error("failed to get grpc host", zap.String("grpc host", GrpcHost))
		return nil, fmt.Errorf("env %s is empty", GrpcHost)
	}

	port := os.Getenv(GrpcPort)
	if len(port) == 0 {
		logger.Error("failed to get grpc port", zap.String("grpc port", GrpcPort))
		return nil, fmt.Errorf("env %s is empty", GrpcPort)
	}

	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}
