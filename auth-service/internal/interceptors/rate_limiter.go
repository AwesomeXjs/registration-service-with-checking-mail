package interceptors

import (
	"context"

	ratelimiter "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/rate_limiter"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RateLimiter struct {
	rateLimiter *ratelimiter.TokenBucketLimiter
}

func NewRateLimitInterceptor(rateLimiter *ratelimiter.TokenBucketLimiter) *RateLimiter {
	return &RateLimiter{rateLimiter: rateLimiter}
}

func (r *RateLimiter) Unary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if !r.rateLimiter.Allow() {
		return nil, status.Error(codes.ResourceExhausted, "rate limit exceeded")
	}

	return handler(ctx, req)
}
