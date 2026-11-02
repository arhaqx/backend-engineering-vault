package consistent_hashing

import "testing"

func TestSolutionV9(t *testing.T) {
	s := NewV9()
	input := []int{1, 2, 3}
	res := s.Compute(input)
	expected := 6 * 9
	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
