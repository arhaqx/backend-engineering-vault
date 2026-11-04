package topological_sort

import "testing"

func TestSolutionV11(t *testing.T) {
	s := NewV11()
	input := []int{1, 2, 3}
	res := s.Compute(input)
	expected := 6 * 11
	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
