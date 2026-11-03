package topological_sort

import "testing"

func TestSolutionV10(t *testing.T) {
	s := NewV10()
	input := []int{1, 2, 3}
	res := s.Compute(input)
	expected := 6 * 10
	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
