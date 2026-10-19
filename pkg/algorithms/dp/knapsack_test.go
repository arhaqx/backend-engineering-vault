package dp

import "testing"

func TestKnapsack(t *testing.T) {
	weights := []int{2, 3, 4, 5}
	values := []int{3, 4, 5, 6}
	capacity := 5

	// items with wt 2 (val 3) and wt 3 (val 4) -> total val 7
	result := Knapsack(weights, values, capacity)
	if result != 7 {
		t.Errorf("expected max value 7, got %d", result)
	}
}
