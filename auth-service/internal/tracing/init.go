package tracing

import (
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"
	"github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
)

// Init initializes the global tracer for the application using Jaeger as the tracing backend.
// It accepts the service name as a parameter, which will be used to identify the service in tracing data.
// The tracer is configured with a constant sampler that always samples all requests.
// If the tracer initialization fails, the function logs the error and terminates the application.
func Init(serviceName string) {

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
