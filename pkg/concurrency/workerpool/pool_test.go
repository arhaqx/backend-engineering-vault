package workerpool

import (
	"context"
	"testing"
	"time"
)

func TestWorkerPool(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	pool := New[int, int](3, 10)
	pool.Start(ctx)

	go func() {
		for i := 1; i <= 5; i++ {
			taskID := string(rune('A' - 1 + i))
			pool.Submit(Task[int, int]{
				ID:   taskID,
				Data: i,
				Exec: func(ctx context.Context, data int) (int, error) {
					return data * 2, nil
				},
			})
		}
		pool.Stop()
	}()

	results := make(map[string]int)
	for res := range pool.Results() {
		if res.Err != nil {
			t.Errorf("task failed: %v", res.Err)
		}
		results[res.TaskID] = res.Output
	}

	if len(results) != 5 {
		t.Fatalf("expected 5 results, got %d", len(results))
	}
}
