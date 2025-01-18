package metrics

import (
	"fmt"
	"net"
	"os"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/pkg/logger"

	"go.uber.org/zap"
)

// Constants for Prometheus environment variable names
const (
	PrometheusPortEnvName = "PROMETHEUS_HTTP_PORT"
	PrometheusHostEnvName = "PROMETHEUS_HTTP_HOST"
)

// PrometheusConfig interface defines the Address method
type PrometheusConfig interface {
	Address() string
}

// PromoConfig struct holds the Prometheus host and port
type PromoConfig struct {
	host string
	port string
}

// NewPrometheusConfig initializes a new Prometheus configuration
func NewPrometheusConfig() (PrometheusConfig, error) {

	const mark = "app.NewPrometheusConfig"

	// Get the Prometheus port and host from environment variables
	port := os.Getenv(PrometheusPortEnvName)
	host := os.Getenv(PrometheusHostEnvName)
	if len(port) == 0 || len(host) == 0 {
		// Log error if either the port or host is missing
		logger.Error("failed to get metrics host",
			mark,
			zap.String("metrics host", PrometheusHostEnvName),
			zap.String("metrics port", PrometheusPortEnvName))
		return nil, fmt.Errorf("PROMETHEUS_HTTP_PORT or PROMETHEUS_HTTP_HOST is not set")
	}

	// Return the Prometheus config with host and port
	return &PromoConfig{
		host: host,
		port: port,
	}, nil
}

// Address returns the full address for the Prometheus config
func (m *PromoConfig) Address() string {
	return net.JoinHostPort(m.host, m.port)
}
