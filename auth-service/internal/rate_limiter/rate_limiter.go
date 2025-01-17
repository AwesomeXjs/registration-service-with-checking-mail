package ratelimiter

import (
	"context"
	"time"
)

// TokenBucketLimiter implements a rate limiter using the token bucket algorithm.
type TokenBucketLimiter struct {
	tokenBucket chan struct{} // Channel to store available tokens.
}

// NewTokenBucketLimiter creates a new TokenBucketLimiter with the specified limit and period.
// Tokens are replenished at regular intervals to maintain the defined rate limit.
func NewTokenBucketLimiter(ctx context.Context, limit int, period time.Duration) *TokenBucketLimiter {
	limiter := &TokenBucketLimiter{
		tokenBucket: make(chan struct{}, limit),
	}

	// Fill the token bucket initially with the maximum allowed tokens.
	for i := 0; i < limit; i++ {
		limiter.tokenBucket <- struct{}{}
	}

	// Calculate the replenishment interval and start periodic token replenishment.
	replenishmentInterval := period.Nanoseconds() / int64(limit)
	go limiter.startPeriodicReplenishment(ctx, time.Duration(replenishmentInterval))

	return limiter
}

// startPeriodicReplenishment periodically adds tokens to the bucket at the specified interval.
func (l *TokenBucketLimiter) startPeriodicReplenishment(ctx context.Context, replenishmentInterval time.Duration) {
	ticker := time.NewTicker(replenishmentInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done(): // Stop replenishment when the context is canceled.
			return
		case <-ticker.C: // Add a token to the bucket at each tick.
			l.tokenBucket <- struct{}{}
		}
	}
}

// Allow attempts to consume a token from the bucket and returns true if successful.
// Returns false if the bucket is empty, indicating the limit has been reached.
func (l *TokenBucketLimiter) Allow() bool {
	select {
	case <-l.tokenBucket:
		return true
	default:
		return false
	}
}
