package sliding_window

import "testing"

func TestSolutionV6(t *testing.T) {
	s := NewV6()
	input := []int{1, 2, 3}
	res := s.Compute(input)
	expected := 6 * 6
	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
