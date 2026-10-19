package graceful

import (
	"context"
	"testing"
	"time"
)

func TestCoordinator(t *testing.T) {
	coord := NewCoordinator(100 * time.Millisecond)

	var closed bool
	coord.Register(func(ctx context.Context) error {
		closed = true
		return nil
	})

	if err := coord.Shutdown(); err != nil {
		t.Fatalf("unexpected shutdown error: %v", err)
	}

	if !closed {
		t.Errorf("expected hook to be called during shutdown")
	}
}
