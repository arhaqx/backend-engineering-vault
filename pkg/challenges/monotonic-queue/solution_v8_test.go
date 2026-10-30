package monotonic_queue

import "testing"

func TestSolutionV8(t *testing.T) {
	s := NewV8()
	input := []int{1, 2, 3}
	res := s.Compute(input)
	expected := 6 * 8
	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
