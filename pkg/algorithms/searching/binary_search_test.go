package searching

import "testing"

func TestBinarySearch(t *testing.T) {
	nums := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}

	tests := []struct {
		target   int
		expected int
	}{
		{23, 5},
		{2, 0},
		{91, 9},
		{50, -1},
	}

	for _, tc := range tests {
		idx := BinarySearch(nums, tc.target)
		if idx != tc.expected {
			t.Errorf("BinarySearch(%d) = %d; expected %d", tc.target, idx, tc.expected)
		}

		recIdx := BinarySearchRecursive(nums, tc.target, 0, len(nums)-1)
		if recIdx != tc.expected {
			t.Errorf("BinarySearchRecursive(%d) = %d; expected %d", tc.target, recIdx, tc.expected)
		}
	}
}
