package two_phase_commit

import "testing"

func TestSolutionV17(t *testing.T) {
	s := NewV17()
	input := []int{1, 2, 3}
	res := s.Compute(input)
	expected := 6 * 17
	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
