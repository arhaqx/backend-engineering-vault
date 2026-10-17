package circuitbreaker

import (
	"errors"
	"sync"
	"time"
)

type State int

const (
	StateClosed State = iota
	StateHalfOpen
	StateOpen
)

var (
	ErrCircuitOpen = errors.New("circuit breaker is open; request rejected")
)

// CircuitBreaker prevents cascading failures when downstream services fail.
type CircuitBreaker struct {
	mu           sync.RWMutex
	state        State
	maxFailures  int
	failureCount int
	timeout      time.Duration
	lastStateChange time.Time
}

// New creates a new CircuitBreaker.
func New(maxFailures int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:        StateClosed,
		maxFailures:  maxFailures,
		timeout:      timeout,
		lastStateChange: time.Now(),
	}
}

// Execute wraps an operation with circuit breaking logic.
func (cb *CircuitBreaker) Execute(fn func() error) error {
	cb.mu.Lock()
	now := time.Now()

	// Check if open state has timed out -> switch to half-open
	if cb.state == StateOpen && now.Sub(cb.lastStateChange) > cb.timeout {
		cb.state = StateHalfOpen
		cb.lastStateChange = now
	}

	if cb.state == StateOpen {
		cb.mu.Unlock()
		return ErrCircuitOpen
	}
	cb.mu.Unlock()

	err := fn()

	cb.mu.Lock()
	defer cb.mu.Unlock()

	if err != nil {
		cb.failureCount++
		if cb.failureCount >= cb.maxFailures {
			cb.state = StateOpen
			cb.lastStateChange = time.Now()
		}
		return err
	}

	// Success resets breaker
	cb.failureCount = 0
	cb.state = StateClosed
	return nil
}

// State returns current breaker state.
func (cb *CircuitBreaker) State() State {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.state
}
