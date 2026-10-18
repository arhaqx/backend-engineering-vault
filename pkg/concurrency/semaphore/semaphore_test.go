package semaphore

import (
	"context"
	"testing"
	"time"
)

func TestSemaphore(t *testing.T) {
	sem := New(2)
	ctx := context.Background()

	if err := sem.Acquire(ctx, 2); err != nil {
		t.Fatalf("failed acquiring initial weight: %v", err)
	}

	ctxTimeout, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()

	if err := sem.Acquire(ctxTimeout, 1); err == nil {
		t.Fatalf("expected acquire to block and time out")
	}

	sem.Release(2)
	if err := sem.Acquire(ctx, 1); err != nil {
		t.Fatalf("failed acquiring after release: %v", err)
	}
}
