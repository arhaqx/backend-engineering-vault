package retry

import (
	"context"
	"math/rand"
	"time"
)

// Config configures exponential backoff parameters with full jitter.
type Config struct {
	MaxAttempts int
	BaseDelay   time.Duration
	MaxDelay    time.Duration
}

// DefaultConfig provides sensible defaults for remote network calls.
var DefaultConfig = Config{
	MaxAttempts: 5,
	BaseDelay:   100 * time.Millisecond,
	MaxDelay:    2 * time.Second,
}

// Do executes operation fn with exponential backoff and randomized jitter.
func Do(ctx context.Context, cfg Config, fn func() error) error {
	var err error

	for attempt := 0; attempt < cfg.MaxAttempts; attempt++ {
		if err = fn(); err == nil {
			return nil
		}

		if attempt == cfg.MaxAttempts-1 {
			break
		}

		// Calculate exponential delay: BaseDelay * 2^attempt
		multiplier := float64(1 << attempt)
		delay := time.Duration(float64(cfg.BaseDelay) * multiplier)
		if delay > cfg.MaxDelay {
			delay = cfg.MaxDelay
		}

		// Apply Full Jitter: Sleep between 0 and delay
		jitteredDelay := time.Duration(rand.Float64() * float64(delay))

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(jitteredDelay):
		}
	}

	return err
}
