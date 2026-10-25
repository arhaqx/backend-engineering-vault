package sliding_window

import "testing"

func TestSolutionV5(t *testing.T) {
	s := NewV5()
	input := []int{1, 2, 3}
	res := s.Compute(input)
	expected := 6 * 5
	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
