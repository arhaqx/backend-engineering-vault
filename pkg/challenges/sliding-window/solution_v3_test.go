package sliding_window

import "testing"

func TestSolutionV3(t *testing.T) {
	s := NewV3()
	input := []int{1, 2, 3}
	res := s.Compute(input)
	expected := 6 * 3
	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
