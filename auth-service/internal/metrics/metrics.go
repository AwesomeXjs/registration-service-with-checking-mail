package metrics

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Constants for Prometheus metrics configuration
const (
	namespace = "auth_service_space"
	appName   = "auth_service"
	subsystem = "grpc"
)

// Metrics structure contains all the Prometheus metrics for the service
type Metrics struct {
	requestCounter        prometheus.Counter       // Counter for total requests
	registrationCounter   prometheus.Counter       // Counter for total registrations
	verificationCounter   prometheus.Counter       // Counter for total verifications
	responseCounter       *prometheus.CounterVec   // Counter for responses with status and method labels
	histogramResponseTime *prometheus.HistogramVec // Histogram for response times
}

// Global variable to hold the metrics object
var metrics *Metrics

// Init  Initializes Prometheus metrics
func Init(_ context.Context) error {
	metrics = &Metrics{
		requestCounter: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      appName + "_requests_total",
			Help:      "Total number of requests.",
		}),
		registrationCounter: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      appName + "_registration_total",
			Help:      "Total number of registrations.",
		}),
		verificationCounter: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      appName + "_verification_total",
			Help:      "Total number of verifications.",
		}),
		responseCounter: promauto.NewCounterVec(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      appName + "_responses_total",
			Help:      "Total number of responses with status and method labels.",
		}, []string{"status", "method"}),
		histogramResponseTime: promauto.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      appName + "_response_time_second",
			Help:      "Response time in seconds.",
			Buckets:   prometheus.ExponentialBuckets(0.001, 2, 16),
		}, []string{"status"}),
	}

	return nil
}

// IncRequestCounter Increments the total request counter
func IncRequestCounter() {
	metrics.requestCounter.Inc()
}

// IncRegistrationCounter Increments the total registration counter
func IncRegistrationCounter() {
	metrics.registrationCounter.Inc()
}

// IncVerificationCounter Increments the total verification counter
func IncVerificationCounter() {
	metrics.verificationCounter.Inc()
}

// IncResponseCounter Increments the response counter with status and method labels
func IncResponseCounter(status, method string) {
	metrics.responseCounter.WithLabelValues(status, method).Inc()
}

// ObserveResponseTime Observes the response time and records it in the histogram
func ObserveResponseTime(status string, time float64) {
	metrics.histogramResponseTime.WithLabelValues(status).Observe(time)
}
