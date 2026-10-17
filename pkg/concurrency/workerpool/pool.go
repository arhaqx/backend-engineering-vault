package workerpool

import (
	"context"
	"sync"
)

// Task represents an executable unit of work.
type Task[T any, R any] struct {
	ID    string
	Data  T
	Exec  func(ctx context.Context, data T) (R, error)
}

// Result represents the outcome of a Task execution.
type Result[R any] struct {
	TaskID string
	Output R
	Err    error
}

// Pool manages a pool of concurrent workers.
type Pool[T any, R any] struct {
	numWorkers int
	tasksChan  chan Task[T, R]
	resultsChan chan Result[R]
	wg         sync.WaitGroup
}

// New creates a new WorkerPool with a specified worker count and buffer size.
func New[T any, R any](numWorkers int, bufferSize int) *Pool[T, R] {
	return &Pool[T, R]{
		numWorkers:  numWorkers,
		tasksChan:   make(chan Task[T, R], bufferSize),
		resultsChan: make(chan Result[R], bufferSize),
	}
}

// Start spawns the worker goroutines.
func (p *Pool[T, R]) Start(ctx context.Context) {
	for i := 0; i < p.numWorkers; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case task, ok := <-p.tasksChan:
					if !ok {
						return
					}
					out, err := task.Exec(ctx, task.Data)
					p.resultsChan <- Result[R]{
						TaskID: task.ID,
						Output: out,
						Err:    err,
					}
				}
			}
		}()
	}
}

// Submit enqueues a new task for processing.
func (p *Pool[T, R]) Submit(task Task[T, R]) {
	p.tasksChan <- task
}

// Stop closes the task channel and waits for all workers to finish.
func (p *Pool[T, R]) Stop() {
	close(p.tasksChan)
	p.wg.Wait()
	close(p.resultsChan)
}

// Results returns the output channel.
func (p *Pool[T, R]) Results() <-chan Result[R] {
	return p.resultsChan
}
