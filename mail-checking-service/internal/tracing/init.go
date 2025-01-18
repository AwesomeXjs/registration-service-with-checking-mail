package tracing

import (
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/pkg/logger"
	"github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
)

// Init initializes the global tracer for the service using Jaeger configuration.
// It sets up a constant sampler to ensure that all traces are captured.
// If tracer initialization fails, the application terminates with a fatal error.
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
