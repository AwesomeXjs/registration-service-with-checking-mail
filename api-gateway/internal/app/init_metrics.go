package app

import (
	"context"
	"log"
	"net/http"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/metrics"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/pkg/logger"
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
