package interceptors

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"

	ratelimiter "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/rate_limiter"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RateLimiter provides a gRPC interceptor to enforce rate limiting on requests.
type RateLimiter struct {
	rateLimiter *ratelimiter.TokenBucketLimiter // Token bucket rate limiter instance.
}

// NewRateLimitInterceptor creates a new RateLimiter interceptor with the provided rate limiter.
func NewRateLimitInterceptor(rateLimiter *ratelimiter.TokenBucketLimiter) *RateLimiter {
	return &RateLimiter{rateLimiter: rateLimiter}
}

// Unary is a gRPC unary interceptor that applies rate limiting to incoming requests.
// Returns a rate limit exceeded error if the request exceeds the allowed limit.
func (r *RateLimiter) Unary(ctx context.Context,
	req interface{},
	_ *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	const mark = "Interceptors.RateLimitInterceptor"

	// Check if the rate limiter allows the request.
	if !r.rateLimiter.Allow() {
		logger.Warn("rate limit exceeded", mark)
		return nil, status.Error(codes.ResourceExhausted, "rate limit exceeded")
	}

	// Proceed with the request if allowed.
	return handler(ctx, req)
}
