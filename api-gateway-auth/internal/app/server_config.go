package app

import (
	"fmt"
	"net"
	"os"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/logger"
	"go.uber.org/zap"
)

// Define constants for the environment variable names
var (
	// HTTPHost refers to the environment variable for the HTTP server host.
	HTTPHost = "HTTP_HOST"

	// HTTPPort refers to the environment variable for the HTTP server port.
	HTTPPort = "HTTP_PORT"
)

// IHTTPConfig interface defines a method to get the address.
type IHTTPConfig interface {
	// Address returns the full address (host and port) as a string.
	Address() string
}

// HTTPConfig is a struct that holds the host and port information.
type HTTPConfig struct {
	host string // The host (e.g., "localhost")
	port string // The port (e.g., "8080")
}

// NewHTTPConfig creates and returns a new instance of HttpConfig.
// It reads the HTTP host and port from environment variables and returns an error if either is missing.
func NewHTTPConfig() (*HTTPConfig, error) {
	// Retrieve the HTTP host from environment variables.
	host := os.Getenv(HTTPHost)
	if len(host) == 0 {
		// Log an error if the host is not set.
		logger.Error("failed to get http host", zap.String("http host", HTTPHost))
		return nil, fmt.Errorf("env %s is empty", HTTPHost)
	}

	// Retrieve the HTTP port from environment variables.
	port := os.Getenv(HTTPPort)
	if len(port) == 0 {
		// Log an error if the port is not set.
		logger.Error("failed to get http port", zap.String("http port", HTTPPort))
		return nil, fmt.Errorf("env %s is empty", HTTPPort)
	}

	// Return a new HttpConfig instance with the host and port values.
	return &HTTPConfig{host: host, port: port}, nil
}

// Address returns the full address as a string by combining the host and port.
func (h *HTTPConfig) Address() string {
	return net.JoinHostPort(h.host, h.port)
}
