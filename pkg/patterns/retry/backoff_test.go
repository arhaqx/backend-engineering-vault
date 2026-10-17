package retry

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestExponentialBackoff(t *testing.T) {
	ctx := context.Background()
	cfg := Config{
		MaxAttempts: 3,
		BaseDelay:   10 * time.Millisecond,
		MaxDelay:    50 * time.Millisecond,
	}

	attempts := 0
	expectedErr := errors.New("temporary service error")

	err := Do(ctx, cfg, func() error {
		attempts++
		if attempts < 3 {
			return expectedErr
		}
		return nil
	})

	if err != nil {
		t.Fatalf("expected operation to eventually succeed, got: %v", err)
	}

	if attempts != 3 {
		t.Fatalf("expected 3 attempts, got %d", attempts)
	}
}
