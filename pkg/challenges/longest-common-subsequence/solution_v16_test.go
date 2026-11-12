package longest_common_subsequence

import "testing"

func TestSolutionV16(t *testing.T) {
	s := NewV16()
	input := []int{1, 2, 3}
	res := s.Compute(input)
	expected := 6 * 16
	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
