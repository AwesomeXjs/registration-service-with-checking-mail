package mail_client

import (
	"fmt"
	"net"
	"os"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/logger"
	"go.uber.org/zap"
)

var (
	// AuthHost defines the environment variable key for the gRPC authentication service host.
	AuthHost = "GRPC_MAIL_HOST"

	// AuthPort defines the environment variable key for the gRPC authentication service port.
	AuthPort = "GRPC_MAIL_PORT"
)

// IMailClientConfig defines the interface for an AuthClient.
type IMailClientConfig interface {
	Address() string
}

// MailClientConfig represents a client for interacting with the Auth service.
type MailClientConfig struct {
	Host string // Host for  the Auth service
	Port string // Port for the Auth service
}

// NewMailClient creates a new AuthClient from environment variables.
func NewMailClient() (IMailClientConfig, error) {
	// Retrieve the Auth service host from environment variables.
	host := os.Getenv(AuthHost)
	if len(host) == 0 {
		logger.Error("failed to get mail host", zap.String("mail host", AuthHost))
		return nil, fmt.Errorf("env %s is empty", AuthHost)
	}

	// Retrieve the Auth service port from environment variables.
	port := os.Getenv(AuthPort)
	if len(port) == 0 {
		logger.Error("failed to get mail port", zap.String("mail port", AuthPort))
		return nil, fmt.Errorf("env %s is empty", AuthPort)
	}

	// Return a new AuthClient instance.
	return &MailClientConfig{
		Host: host,
		Port: port,
	}, nil
}

// Address returns the full address of the Auth service (host:port).
func (a *MailClientConfig) Address() string {
	return net.JoinHostPort(a.Host, a.Port)
}
