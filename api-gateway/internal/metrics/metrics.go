package metrics

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	// Namespace for all metrics related to the API Gateway
	namespace = "api_gateway_space"
	// Application name for metrics
	appName = "api_gateway"
	// Subsystem for REST API
	subsystem = "rest"
)

// Metrics struct contains all the metrics for the API Gateway.
type Metrics struct {
	// requestCounter tracks the total number of requests made to the server.
	requestCounter prometheus.Counter
	// responseCounter tracks the number of responses by status and method.
	responseCounter *prometheus.CounterVec
	// histogramResponseTime tracks the response time for each status of the request.
	histogramResponseTime *prometheus.HistogramVec
}

// metrics variable holds the metrics for the current application.
var metrics *Metrics

// Init initializes the metrics for the API Gateway.
// It creates counters and histograms to track requests, responses, and response time.
func Init(_ context.Context) error {
	metrics = &Metrics{
		requestCounter: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      appName + "_requests_total",
			Help:      "Total number of requests to the server.",
		}),
		responseCounter: promauto.NewCounterVec(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      appName + "_responses_total",
			Help:      "Total number of responses from the server.",
		}, []string{"status", "method"}),
		histogramResponseTime: promauto.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      appName + "_response_time_second",
			Help:      "Request execution time.",
			Buckets:   prometheus.ExponentialBuckets(0.001, 2, 16),
		}, []string{"status"}),
	}

	return nil
}

// IncRequestCounter increments the request counter by 1.
func IncRequestCounter() {
	metrics.requestCounter.Inc()
}

// IncResponseCounter increments the response counter by 1 for a specific status and method.
func IncResponseCounter(status, method string) {
	metrics.responseCounter.WithLabelValues(status, method).Inc()
}

// ObserveResponseTime records the response time for a specific status.
func ObserveResponseTime(status string, time float64) {
	metrics.histogramResponseTime.WithLabelValues(status).Observe(time)
}
