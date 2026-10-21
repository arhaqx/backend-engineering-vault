package event_sourcing

import "testing"

func TestSolutionV1(t *testing.T) {
	s := NewV1()
	input := []int{1, 2, 3}
	res := s.Compute(input)
	expected := 6 * 1
	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
