package pipeline

import (
	"sync"
)

// Generator emits values from a slice into a receive-only channel.
func Generator[T any](done <-chan struct{}, values ...T) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for _, v := range values {
			select {
			case <-done:
				return
			case out <- v:
			}
		}
	}()
	return out
}

// Stage transforms values from in channel using a worker mapping function.
func Stage[T any, R any](done <-chan struct{}, in <-chan T, transform func(T) R) <-chan R {
	out := make(chan R)
	go func() {
		defer close(out)
		for val := range in {
			select {
			case <-done:
				return
			case out <- transform(val):
			}
		}
	}()
	return out
}

// FanIn multiplexes multiple channels into a single consolidated output channel.
func FanIn[T any](done <-chan struct{}, channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	multiplexed := make(chan T)

	multiplex := func(c <-chan T) {
		defer wg.Done()
		for val := range c {
			select {
			case <-done:
				return
			case multiplexed <- val:
			}
		}
	}

	wg.Add(len(channels))
	for _, ch := range channels {
		go multiplex(ch)
	}

	go func() {
		wg.Wait()
		close(multiplexed)
	}()

	return multiplexed
}
