package searching

import "cmp"

// BinarySearch performs iterative binary search on a sorted slice.
// Returns the index of the element if found, or -1 if not found.
func BinarySearch[T cmp.Ordered](arr []T, target T) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

// BinarySearchRecursive performs recursive binary search on a sorted slice.
func BinarySearchRecursive[T cmp.Ordered](arr []T, target T, low, high int) int {
	if low > high {
		return -1
	}

	mid := low + (high-low)/2
	if arr[mid] == target {
		return mid
	}

	if arr[mid] > target {
		return BinarySearchRecursive(arr, target, low, mid-1)
	}

	return BinarySearchRecursive(arr, target, mid+1, high)
}
