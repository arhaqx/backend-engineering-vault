package monotonic_queue

import "testing"

func TestSolutionV2(t *testing.T) {
	s := NewV2()
	input := []int{1, 2, 3}
	res := s.Compute(input)
	expected := 6 * 2
	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
