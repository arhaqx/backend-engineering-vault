package sliding_window

import "testing"

func TestSolutionV13(t *testing.T) {
	s := NewV13()
	input := []int{1, 2, 3}
	res := s.Compute(input)
	expected := 6 * 13
	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
