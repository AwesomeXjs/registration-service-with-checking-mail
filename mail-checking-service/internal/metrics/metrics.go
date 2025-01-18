package metrics

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Constants for Prometheus metric names and namespaces
const (
	namespace = "mail_service_space"
	appName   = "mail_service"
	subsystem = "grpc"
)

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
		// Counter for total requests to the server
		requestCounter: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      appName + "_requests_total",
			Help:      "Количество запросов к серверу.",
		}),
		// Counter for successful verification events
		successVerificationCounter: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      appName + "_success_verification_total",
			Help:      "Количество успешных проверок.",
		}),
		// Counter for events from Kafka
		kafkaEventCounter: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      appName + "_kafka_events_total",
			Help:      "Количество событий из Kafka.",
		}),
		// Counter for responses from the server, categorized by status and method
		responseCounter: promauto.NewCounterVec(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      appName + "_responses_total",
			Help:      "Количество ответов от сервера.",
		}, []string{"status", "method"}),
		// Histogram for tracking response time with exponential buckets
		histogramResponseTime: promauto.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      appName + "_response_time_second",
			Help:      "Время выполнения запроса.",
			Buckets:   prometheus.ExponentialBuckets(0.001, 2, 16),
		}, []string{"status"}),
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
