package metrics

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

// Constants for Prometheus metric names and namespaces

// Metrics struct holds all the Prometheus metrics for the service
type Metrics struct {
	requestCounter             prometheus.Counter
	successVerificationCounter prometheus.Counter
	kafkaEventCounter          prometheus.Counter
	responseCounter            *prometheus.CounterVec
	histogramResponseTime      *prometheus.HistogramVec
}

// Global variable to store metrics instance
var metrics *Metrics

// Init initializes the Prometheus metrics for the mail service
func Init(_ context.Context) error {
	metrics = &Metrics{
		requestCounter:             NewCounter("requests_total", "Total number of requests."),
		successVerificationCounter: NewCounter("success_verifications_total", "Total number of successful verifications."),
		kafkaEventCounter:          NewCounter("kafka_events_total", "Total number of events from Kafka."),
		responseCounter:            NewCounterVec("responses_total", "Total number of responses.", []string{"status", "method"}),
		histogramResponseTime: NewHistogramVec("response_time_second",
			"Request execution time.",
			prometheus.ExponentialBuckets(0.001, 2, 16),
			[]string{"status"}),
	}

	return nil
}

// IncRequestCounter increments the request counter
func IncRequestCounter() {
	metrics.requestCounter.Inc()
}

// IncSuccessVerificationCounter increments the successful verification counter
func IncSuccessVerificationCounter() {
	metrics.successVerificationCounter.Inc()
}

// IncKafkaEventCounter increments the Kafka event counter
func IncKafkaEventCounter() {
	metrics.kafkaEventCounter.Inc()
}

// IncResponseCounter increments the response counter, categorized by status and method
func IncResponseCounter(status, method string) {
	metrics.responseCounter.WithLabelValues(status, method).Inc()
}

// ObserveResponseTime records the response time for a given status
func ObserveResponseTime(status string, time float64) {
	metrics.histogramResponseTime.WithLabelValues(status).Observe(time)
}
