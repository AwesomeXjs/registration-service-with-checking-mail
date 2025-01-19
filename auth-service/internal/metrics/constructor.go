package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Constants for Prometheus metric identification.
const (
	namespace = "auth_service_space" // Metric namespace.
	appName   = "auth_service"       // Application name.
	subsystem = "grpc"               // Application subsystem.
)

// NewCounter creates a Prometheus counter metric.
// name: Metric name. description: Metric description.
func NewCounter(name, description string) prometheus.Counter {
	return promauto.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      appName + "_" + name,
		Help:      description,
	})
}

// NewCounterVec creates a Prometheus counter vector.
// name: Metric name. description: Metric description. labels: Metric labels.
func NewCounterVec(name, description string, labels []string) *prometheus.CounterVec {
	return promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      appName + "_" + name,
		Help:      description,
	}, labels)
}

// NewHistogramVec creates a Prometheus histogram vector.
// name: Metric name. description: Metric description.
// buckets: Histogram buckets. labels: Metric labels.
func NewHistogramVec(
	name, description string, buckets []float64, labels []string,
) *prometheus.HistogramVec {
	return promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      appName + "_" + name,
		Help:      description,
		Buckets:   buckets,
	}, labels)
}
