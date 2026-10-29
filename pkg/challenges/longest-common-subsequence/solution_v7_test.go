package longest_common_subsequence

import "testing"

func TestSolutionV7(t *testing.T) {
	s := NewV7()
	input := []int{1, 2, 3}
	res := s.Compute(input)
	expected := 6 * 7
	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
