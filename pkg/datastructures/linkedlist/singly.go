package linkedlist

// Node represents an element in the singly linked list.
type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

// SinglyLinkedList represents a generic singly linked list.
type SinglyLinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

// NewSingly creates a new empty SinglyLinkedList.
func NewSingly[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

// Append adds a new value to the end of the list.
func (l *SinglyLinkedList[T]) Append(val T) {
	node := &Node[T]{Value: val}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.Next = node
		l.tail = node
	}
	l.size++
}

// Prepend adds a new value to the beginning of the list.
func (l *SinglyLinkedList[T]) Prepend(val T) {
	node := &Node[T]{Value: val, Next: l.head}
	l.head = node
	if l.tail == nil {
		l.tail = node
	}
	l.size++
}

// Size returns current number of elements.
func (l *SinglyLinkedList[T]) Size() int {
	return l.size
}

// Find searches for a value and returns true if found.
func (l *SinglyLinkedList[T]) Find(val T) bool {
	curr := l.head
	for curr != nil {
		if curr.Value == val {
			return true
		}
		curr = curr.Next
	}
	return false
}

// ToSlice converts linked list values into a slice.
func (l *SinglyLinkedList[T]) ToSlice() []T {
	res := make([]T, 0, l.size)
	curr := l.head
	for curr != nil {
		res = append(res, curr.Value)
		curr = curr.Next
	}
	return res
}
