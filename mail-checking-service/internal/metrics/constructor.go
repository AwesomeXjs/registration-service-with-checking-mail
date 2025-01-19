package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Namespace, app name, and subsystem for metrics.
const (
	namespace = "mail_service_space" // Metric namespace.
	appName   = "mail_service"       // Application name.
	subsystem = "grpc"               // Application subsystem.
)

// NewCounter creates and registers a Prometheus counter.
// name: Metric name. description: Metric description.
func NewCounter(name, description string) prometheus.Counter {
	return promauto.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      appName + "_" + name,
		Help:      description,
	})
}

// NewCounterVec creates and registers a counter vector metric.
// name: Metric name. description: Metric description. labels: Metric labels.
func NewCounterVec(name, description string, labels []string) *prometheus.CounterVec {
	return promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      appName + "_" + name,
		Help:      description,
	}, labels)
}

// NewHistogramVec creates and registers a histogram vector.
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
