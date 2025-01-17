package linkedlist

import "errors"

type node[T any] struct {
	Value T
	Next  *node[T]
}

type LinkedList[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (ll *LinkedList[T]) Add(value T) {
	node := &node[T]{Value: value}

	if ll.head == nil {
		ll.head = node
		ll.tail = node
	} else {
		ll.tail.Next = node
		ll.tail = node
	}

	ll.size++
}

func (ll *LinkedList[T]) Set(index int, value T) bool {
	if index < 0 || index >= ll.size {
		return false
	}

	current := ll.head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	current.Value = value
	return true
}

func (ll *LinkedList[T]) Remove(index int) bool {
	if index < 0 || index >= ll.size {
		return false
	}

	// Remove head
	if index == 0 {
		ll.head = ll.head.Next
		if ll.head == nil {
			ll.tail = nil
		}
		ll.size--
		return true
	}

	current := ll.head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	current.Next = current.Next.Next
	if current.Next == nil {
		ll.tail = current
	}
	ll.size--
	return true
}

func (ll *LinkedList[T]) Clear() {
	ll.head = nil
	ll.tail = nil
	ll.size = 0
}

func (ll *LinkedList[T]) Get(index int) (T, error) {
	var value T
	if index < 0 || index >= ll.size {
		return value, errors.New("index out of bounds")
	}

	current := ll.head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Value, nil
}

func (ll *LinkedList[T]) Size() int {
	return ll.size
}
