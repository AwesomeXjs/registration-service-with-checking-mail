package app

import (
	"context"
	"log"
	"net/http"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"
	"github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/metrics"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (a *App) initPrometheus(_ context.Context) error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	a.prometheus = &http.Server{
		Addr:    a.serviceProvider.PrometheusConfig().Address(),
		Handler: mux,
	}
	return nil
}

func (a *App) runPrometheus() error {
	logger.Info("starting prometheus server on "+a.serviceProvider.PrometheusConfig().Address(), "App.app.runPrometheus")
	return a.prometheus.ListenAndServe()
}

func (a *App) initMetrics(ctx context.Context) error {
	err := metrics.Init(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// InitTracing initializes the global tracer for the application using Jaeger as the tracing backend.
// It accepts the service name as a parameter, which will be used to identify the service in tracing data.
// The tracer is configured with a constant sampler that always samples all requests.
// If the tracer initialization fails, the function logs the error and terminates the application.
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
