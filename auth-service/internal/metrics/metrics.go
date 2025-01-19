package metrics

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

// Constants for Prometheus metrics configuration

// Metrics structure contains all the Prometheus metrics for the service
type Metrics struct {
	requestCounter        prometheus.Counter
	registrationCounter   prometheus.Counter
	verificationCounter   prometheus.Counter
	responseCounter       *prometheus.CounterVec
	histogramResponseTime *prometheus.HistogramVec
}

// Global variable to hold the metrics object
var metrics *Metrics

// Init  Initializes Prometheus metrics
func Init(_ context.Context) error {
	metrics = &Metrics{
		requestCounter:      NewCounter("requests_total", "Total number of requests."),
		registrationCounter: NewCounter("registrations_total", "Total number of registrations."),
		verificationCounter: NewCounter("verifications_total", "Total number of verifications."),
		responseCounter:     NewCounterVec("responses_total", "Total number of responses.", []string{"status", "method"}),
		histogramResponseTime: NewHistogramVec("response_time_second",
			"Request execution time.", prometheus.ExponentialBuckets(0.001, 2, 16),
			[]string{"status"}),
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
