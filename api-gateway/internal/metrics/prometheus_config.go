package metrics

import (
	"fmt"
	"net"
	"os"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/pkg/logger"
	"go.uber.org/zap"
)

const (
	// PrometheusPortEnvName is the environment variable name for the Prometheus HTTP port.
	PrometheusPortEnvName = "PROMETHEUS_HTTP_PORT"
	// PrometheusHostEnvName is the environment variable name for the Prometheus HTTP host.
	PrometheusHostEnvName = "PROMETHEUS_HTTP_HOST"
)

// PrometheusConfig is an interface defining the method to retrieve the address for Prometheus metrics.
type PrometheusConfig interface {
	// Address returns the full address of the Prometheus server (host and port).
	Address() string
}

// PromoConfig holds the configuration details for Prometheus server.
type PromoConfig struct {
	// host is the hostname or IP address of the Prometheus server.
	host string
	// port is the port number on which Prometheus server is running.
	port string
}

// NewPrometheusConfig creates a new PrometheusConfig instance by reading the host and port
// from environment variables. Returns an error if either value is missing.
func NewPrometheusConfig() (PrometheusConfig, error) {

	// Mark for logging
	const mark = "app.NewPrometheusConfig"

	// Retrieve Prometheus HTTP port and host from environment variables.
	port := os.Getenv(PrometheusPortEnvName)
	host := os.Getenv(PrometheusHostEnvName)

	// If either the port or host is missing, log an error and return an error.
	if len(port) == 0 || len(host) == 0 {
		logger.Error("failed to get metrics host",
			mark,
			zap.String("metrics host", PrometheusHostEnvName),
			zap.String("metrics port", PrometheusPortEnvName))
		return nil, fmt.Errorf("PROMETHEUS_HTTP_PORT or PROMETHEUS_HTTP_HOST is not set")
	}

	// Return a new instance of PromoConfig with the host and port.
	return &PromoConfig{
		host: host,
		port: port,
	}, nil
}

// Address returns the full address for the Prometheus server by joining the host and port.
func (m *PromoConfig) Address() string {
	return net.JoinHostPort(m.host, m.port)
}
