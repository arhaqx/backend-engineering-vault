package graceful

import (
	"context"
	"sync"
	"time"
)

type Hook func(ctx context.Context) error

type Coordinator struct {
	mu      sync.Mutex
	hooks   []Hook
	timeout time.Duration
}

func NewCoordinator(timeout time.Duration) *Coordinator {
	return &Coordinator{timeout: timeout}
}

func (c *Coordinator) Register(hook Hook) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.hooks = append(c.hooks, hook)
}

func (c *Coordinator) Shutdown() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	var wg sync.WaitGroup
	errChan := make(chan error, len(c.hooks))

	for _, hook := range c.hooks {
		wg.Add(1)
		go func(h Hook) {
			defer wg.Done()
			if err := h(ctx); err != nil {
				errChan <- err
			}
		}(hook)
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return err
		}
	}
	return nil
}
