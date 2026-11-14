package sliding_window

import "testing"

func TestSolutionV18(t *testing.T) {
	s := NewV18()
	input := []int{1, 2, 3}
	res := s.Compute(input)
	expected := 6 * 18
	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
