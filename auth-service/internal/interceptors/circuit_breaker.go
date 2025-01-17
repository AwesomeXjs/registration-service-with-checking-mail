package interceptors

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"
	"go.uber.org/zap"

	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CircuitBreaker provides a gRPC interceptor with a circuit breaker mechanism.
type CircuitBreaker struct {
	cb *gobreaker.CircuitBreaker // Gobreaker circuit breaker instance.
}

// NewCircuitBreaker initializes a new CircuitBreaker with the provided gobreaker instance.
func NewCircuitBreaker(cb *gobreaker.CircuitBreaker) *CircuitBreaker {
	return &CircuitBreaker{cb: cb}
}

// Unary is a gRPC unary interceptor that wraps requests with a circuit breaker.
// If the circuit breaker is in the open state, it returns a service unavailable error.
func (c *CircuitBreaker) Unary(ctx context.Context,
	req interface{},
	_ *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	const mark = "Interceptors.CircuitBreakerInterceptor"

	// Execute the handler function within the circuit breaker.
	res, err := c.cb.Execute(func() (interface{}, error) {
		return handler(ctx, req)
	})
	if err != nil {
		// Return a specific error if the circuit breaker is open.
		if err == gobreaker.ErrOpenState {
			logger.Error("service unavailable", mark, zap.Error(err))
			return nil, status.Error(codes.Unavailable, "service unavailable")
		}
		logger.Error("failed to execute", mark, zap.Error(err))
		return nil, err
	}

	return res, nil
}
