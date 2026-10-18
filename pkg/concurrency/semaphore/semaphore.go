package semaphore

import (
	"context"
	"errors"
	"sync"
)

var ErrNoTickets = errors.New("no tickets available")

// Semaphore controls access to a finite number of shared resources.
type Semaphore struct {
	mu      sync.Mutex
	capacity int64
	current  int64
	waiters  chan struct{}
}

func New(capacity int64) *Semaphore {
	return &Semaphore{
		capacity: capacity,
		waiters:  make(chan struct{}, capacity),
	}
}

func (s *Semaphore) Acquire(ctx context.Context, weight int64) error {
	for {
		s.mu.Lock()
		if s.capacity-s.current >= weight {
			s.current += weight
			s.mu.Unlock()
			return nil
		}
		s.mu.Unlock()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-s.waiters:
		}
	}
}

func (s *Semaphore) Release(weight int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.current -= weight
	select {
	case s.waiters <- struct{}{}:
	default:
	}
}
