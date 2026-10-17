package pipeline

import (
	"sort"
	"testing"
)

func TestFanOutFanInPipeline(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	numbers := []int{1, 2, 3, 4, 5, 6}
	in := Generator(done, numbers...)

	// Fan-out: 2 workers square the numbers concurrently
	square := func(n int) int { return n * n }
	worker1 := Stage(done, in, square)
	worker2 := Stage(done, in, square)

	// Fan-in: combine worker outputs
	out := FanIn(done, worker1, worker2)

	var results []int
	for val := range out {
		results = append(results, val)
	}

	sort.Ints(results)
	expected := []int{1, 4, 9, 16, 25, 36}

	if len(results) != len(expected) {
		t.Fatalf("expected %d elements, got %d", len(expected), len(results))
	}

	for i := range results {
		if results[i] != expected[i] {
			t.Errorf("at index %d: expected %d, got %d", i, expected[i], results[i])
		}
	}
}
