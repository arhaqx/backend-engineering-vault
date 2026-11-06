package lfu_cache

import "testing"

func TestSolutionV12(t *testing.T) {
	s := NewV12()
	input := []int{1, 2, 3}
	res := s.Compute(input)
	expected := 6 * 12
	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
