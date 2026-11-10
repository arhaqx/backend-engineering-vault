package monotonic_queue

import "testing"

func TestSolutionV15(t *testing.T) {
	s := NewV15()
	input := []int{1, 2, 3}
	res := s.Compute(input)
	expected := 6 * 15
	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
