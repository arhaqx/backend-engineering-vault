package singleflight

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestSingleflight(t *testing.T) {
	g := NewGroup()
	var execCount int32

	fn := func() (any, error) {
		atomic.AddInt32(&execCount, 1)
		time.Sleep(20 * time.Millisecond)
		return "result", nil
	}

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res, err := g.Do("heavy-key", fn)
			if err != nil || res != "result" {
				t.Errorf("unexpected output: %v, %v", res, err)
			}
		}()
	}
	wg.Wait()

	if atomic.LoadInt32(&execCount) != 1 {
		t.Fatalf("expected fn to execute exactly once, got %d", execCount)
	}
}
