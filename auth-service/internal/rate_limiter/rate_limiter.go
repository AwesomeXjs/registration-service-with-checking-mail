package ratelimiter

import (
	"context"
	"fmt"
	"time"
)

type TokenBucketLimiter struct {
	tokenBucket chan struct{}
}

func NewTokenBucketLimiter(ctx context.Context, limit int, period time.Duration) *TokenBucketLimiter {
	limiter := &TokenBucketLimiter{
		tokenBucket: make(chan struct{}, limit),
	}

	for i := 0; i < limit; i++ {
		limiter.tokenBucket <- struct{}{}
	}

	replenishmentInterval := period.Nanoseconds() / int64(limit)
	go limiter.startPeriodicReplenishment(ctx, time.Duration(replenishmentInterval))

	return limiter
}

func (l *TokenBucketLimiter) startPeriodicReplenishment(ctx context.Context, replenishmentInterval time.Duration) {
	fmt.Println(replenishmentInterval)
	ticker := time.NewTicker(replenishmentInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			l.tokenBucket <- struct{}{}
		}
	}
}

func (l *TokenBucketLimiter) Allow() bool {
	select {
	case <-l.tokenBucket:
		return true
	default:
		return false
	}
}
