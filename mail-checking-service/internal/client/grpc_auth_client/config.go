package grpc_auth_client

import (
	"fmt"
	"net"
	"os"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/logger"
	"go.uber.org/zap"
)

var (
	// AuthHost defines the environment variable key for the gRPC authentication service host.
	AuthHost = "GRPC_AUTH_HOST"

	// AuthPort defines the environment variable key for the gRPC authentication service port.
	AuthPort = "GRPC_AUTH_PORT"
)

// IAuthClientConfig defines the interface for an AuthClient.
type IAuthClientConfig interface {
	Address() string
}

// AuthClientConfig represents a client for interacting with the Auth service.
type AuthClientConfig struct {
	Host string // Host for the Auth service
	Port string // Port for the Auth service
}

// NewAuthClientConfig creates a new AuthClient from environment variables.
func NewAuthClientConfig() (IAuthClientConfig, error) {
	// Retrieve the Auth service host from environment variables.
	host := os.Getenv(AuthHost)
	if len(host) == 0 {
		logger.Error("failed to get auth host", zap.String("auth host", AuthHost))
		return nil, fmt.Errorf("env %s is empty", AuthHost)
	}

	// Retrieve the Auth service port from environment variables.
	port := os.Getenv(AuthPort)
	if len(port) == 0 {
		logger.Error("failed to get auth port", zap.String("auth port", AuthPort))
		return nil, fmt.Errorf("env %s is empty", AuthPort)
	}

	// Return a new AuthClient instance.
	return &AuthClientConfig{
		Host: host,
		Port: port,
	}, nil
}

// Address returns the full address of the Auth service (host:port).
func (a *AuthClientConfig) Address() string {
	return net.JoinHostPort(a.Host, a.Port)
}
