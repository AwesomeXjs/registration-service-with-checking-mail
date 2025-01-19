package app

import (
	"context"
	"net/http"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/metrics"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/pkg/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
)

const (
	metricsRoute = "/metrics"
)

func (a *App) initPrometheus(_ context.Context) error {
	mux := http.NewServeMux()
	mux.Handle(metricsRoute, promhttp.Handler())

	a.prometheus = &http.Server{
		Addr:    a.serviceProvider.PrometheusConfig().Address(),
		Handler: mux,
	}
	return nil
}

func (a *App) runPrometheus() error {

	const mark = "App.app.runPrometheus"

	logger.Info("starting prometheus server on "+a.serviceProvider.PrometheusConfig().Address(), mark)
	return a.prometheus.ListenAndServe()
}

func (a *App) initMetrics(ctx context.Context) error {

	const mark = "App.app.initMetrics"

	err := metrics.Init(ctx)
	if err != nil {
		logger.Fatal("failed to init metrics", mark, zap.Error(err))
	}
	return nil
}

// InitTracing initializes the global tracer for the service using Jaeger configuration.
// It sets up a constant sampler to ensure that all traces are captured.
// If tracer initialization fails, the application terminates with a fatal error.
func (a *App) InitTracing(serviceName string) {

	const mark = "Tracing.Init"

	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
	}

	_, err := cfg.InitGlobalTracer(serviceName)
	if err != nil {
		logger.Fatal("failed to init tracer", mark, zap.Error(err))
	}
}
