package longest_common_subsequence

import "testing"

func TestSolutionV4(t *testing.T) {
	s := NewV4()
	input := []int{1, 2, 3}
	res := s.Compute(input)
	expected := 6 * 4
	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
