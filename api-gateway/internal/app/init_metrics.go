package app

import (
	"context"
	"net/http"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/metrics"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/pkg/logger"
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

// InitTracing initializes a global Jaeger tracer for the given service name.
// It uses a constant sampler configuration to sample all traces (Param: 1).
// If the tracer initialization fails, the function logs a fatal error
// and terminates the application.
func (a *App) InitTracing(serviceName string) {

	const mark = "Tracing.Init"

	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:          true,
			CollectorEndpoint: "http://jaeger:4318",
		},
	}

	_, err := cfg.InitGlobalTracer(serviceName)
	if err != nil {
		logger.Fatal("failed to init tracer", mark, zap.Error(err))
	}
}
