package monotonic_queue

import "testing"

func TestSolutionV14(t *testing.T) {
	s := NewV14()
	input := []int{1, 2, 3}
	res := s.Compute(input)
	expected := 6 * 14
	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
