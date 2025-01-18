package tracing

import (
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/pkg/logger"
	"github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
)

// Init initializes a global Jaeger tracer for the given service name.
// It uses a constant sampler configuration to sample all traces (Param: 1).
// If the tracer initialization fails, the function logs a fatal error
// and terminates the application.
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
