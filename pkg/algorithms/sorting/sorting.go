package sorting

import "cmp"

// QuickSort sorts a slice of ordered items in ascending order using QuickSort algorithm.
func QuickSort[T cmp.Ordered](arr []T) []T {
	if len(arr) < 2 {
		return arr
	}

	result := make([]T, len(arr))
	copy(result, arr)
	quickSortHelper(result, 0, len(result)-1)
	return result
}

func quickSortHelper[T cmp.Ordered](arr []T, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		quickSortHelper(arr, low, p-1)
		quickSortHelper(arr, p+1, high)
	}
}

func partition[T cmp.Ordered](arr []T, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// MergeSort sorts a slice of ordered items in ascending order using MergeSort.
func MergeSort[T cmp.Ordered](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func merge[T cmp.Ordered](left, right []T) []T {
	result := make([]T, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
