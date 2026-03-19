package main

import (
	"context"

	"golang.org/x/time/rate"
)

// throttler implements global throttling for outgoing messages.
// Unlike the rateLimiter which rejects requests when limits are exceeded,
// the throttler blocks until capacity is available (backpressure).
type throttler struct {
	limiter *rate.Limiter
}

// newThrottler creates a new throttler with the given configuration.
// messagesPerSecond controls the steady-state rate, burst allows short spikes.
func newThrottler(messagesPerSecond float64, burst int) *throttler {
	return &throttler{
		limiter: rate.NewLimiter(rate.Limit(messagesPerSecond), burst),
	}
}

// wait blocks until the throttler allows the request or the context is cancelled.
// Returns nil if allowed, or the context error if cancelled/timed out.
func (t *throttler) wait(ctx context.Context) error {
	return t.limiter.Wait(ctx)
}
