package metrics

import (
	"fmt"
	"net"
	"os"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"
	"go.uber.org/zap"
)

// Environment variable names for Prometheus host and port
const (
	PrometheusPortEnvName = "PROMETHEUS_HTTP_PORT"
	PrometheusHostEnvName = "PROMETHEUS_HTTP_HOST"
)

// PrometheusConfig interface defines the method to get the Prometheus address
type PrometheusConfig interface {
	Address() string
}

// PromoConfig holds the Prometheus host and port configuration
type PromoConfig struct {
	host string // Prometheus host
	port string // Prometheus port
}

// NewPrometheusConfig initializes the Prometheus configuration from environment variables
func NewPrometheusConfig() (PrometheusConfig, error) {

	const mark = "app.NewPrometheusConfig"

	// Get Prometheus host and port from environment variables
	port := os.Getenv(PrometheusPortEnvName)
	host := os.Getenv(PrometheusHostEnvName)
	if len(port) == 0 || len(host) == 0 {
		// Log error and return if required environment variables are not set
		logger.Error("failed to get metrics host",
			mark,
			zap.String("metrics host", PrometheusHostEnvName),
			zap.String("metrics port", PrometheusPortEnvName))
		return nil, fmt.Errorf("PROMETHEUS_HTTP_PORT or PROMETHEUS_HTTP_HOST is not set")
	}

	// Return the new PrometheusConfig instance
	return &PromoConfig{
		host: host,
		port: port,
	}, nil
}

// Address returns the full Prometheus address in the form host:port
func (m *PromoConfig) Address() string {
	return net.JoinHostPort(m.host, m.port)
}
