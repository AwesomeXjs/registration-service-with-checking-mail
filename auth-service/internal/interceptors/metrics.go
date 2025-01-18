package interceptors

import (
	"context"
	"time"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/metrics"
	"google.golang.org/grpc"
)

// MetricsInterceptor - собирает метрики для Prometheus
func MetricsInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	metrics.IncRequestCounter()

	timeStart := time.Now()

	res, err := handler(ctx, req)

	diffTime := time.Since(timeStart)

	if err != nil {
		metrics.IncResponseCounter("error", info.FullMethod)
		metrics.ObserveResponseTime("error", diffTime.Seconds())
	} else {
		metrics.IncResponseCounter("success", info.FullMethod)
		metrics.ObserveResponseTime("success", diffTime.Seconds())
	}

	return res, err
}
