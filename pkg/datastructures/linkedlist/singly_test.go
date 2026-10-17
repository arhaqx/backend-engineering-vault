package linkedlist

import (
	"reflect"
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {
	list := NewSingly[int]()

	list.Append(10)
	list.Append(20)
	list.Prepend(5)

	if list.Size() != 3 {
		t.Fatalf("expected size 3, got %d", list.Size())
	}

	expected := []int{5, 10, 20}
	if !reflect.DeepEqual(list.ToSlice(), expected) {
		t.Fatalf("expected %v, got %v", expected, list.ToSlice())
	}

	if !list.Find(10) {
		t.Errorf("expected to find 10")
	}

	if list.Find(999) {
		t.Errorf("did not expect to find 999")
	}
}
