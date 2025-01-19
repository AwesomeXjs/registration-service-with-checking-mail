package metrics

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
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
		requestCounter:  NewCounter("requests_total", "Total number of requests."),
		responseCounter: NewCounterVec("responses_total", "Total number of responses.", []string{"status", "method"}),
		histogramResponseTime: NewHistogramVec("response_time_second",
			"Request execution time.",
			prometheus.ExponentialBuckets(0.001, 2, 16),
			[]string{"status"}),
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
