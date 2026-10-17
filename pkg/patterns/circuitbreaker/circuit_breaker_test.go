package circuitbreaker

import (
	"errors"
	"testing"
	"time"
)

func TestCircuitBreaker(t *testing.T) {
	cb := New(2, 50*time.Millisecond)

	failFn := func() error {
		return errors.New("downstream error")
	}
	successFn := func() error {
		return nil
	}

	// 1st failure
	_ = cb.Execute(failFn)
	if cb.State() != StateClosed {
		t.Fatalf("expected state Closed after 1 failure")
	}

	// 2nd failure -> Trips breaker to Open
	_ = cb.Execute(failFn)
	if cb.State() != StateOpen {
		t.Fatalf("expected state Open after 2 failures")
	}

	// 3rd call rejected immediately
	err := cb.Execute(successFn)
	if !errors.Is(err, ErrCircuitOpen) {
		t.Fatalf("expected ErrCircuitOpen, got %v", err)
	}

	// Wait for cooldown timeout
	time.Sleep(60 * time.Millisecond)

	// Call succeeds -> resets to Closed
	err = cb.Execute(successFn)
	if err != nil {
		t.Fatalf("expected success, got %v", err)
	}
	if cb.State() != StateClosed {
		t.Fatalf("expected state Closed after recovery")
	}
}
