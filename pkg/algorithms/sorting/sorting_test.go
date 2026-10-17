package sorting

import (
	"reflect"
	"testing"
)

func TestSorting(t *testing.T) {
	input := []int{64, 34, 25, 12, 22, 11, 90}
	expected := []int{11, 12, 22, 25, 34, 64, 90}

	qs := QuickSort(input)
	if !reflect.DeepEqual(qs, expected) {
		t.Errorf("QuickSort failed: got %v, expected %v", qs, expected)
	}

	ms := MergeSort(input)
	if !reflect.DeepEqual(ms, expected) {
		t.Errorf("MergeSort failed: got %v, expected %v", ms, expected)
	}
}
