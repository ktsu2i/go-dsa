package doublylinkedlist

import "errors"

type node[T any] struct {
	value T
	prev  *node[T]
	next  *node[T]
}

type DoublyLinkedList[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

func New[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

func (dll *DoublyLinkedList[T]) Add(value T) {
	node := &node[T]{value: value}

	if dll.head == nil {
		dll.head = node
		dll.tail = node
	} else {
		node.prev = dll.tail
		dll.tail.next = node
		dll.tail = node
	}

	dll.size++
}

func (dll *DoublyLinkedList[T]) Set(index int, value T) bool {
	if index < 0 || index >= dll.size {
		return false
	}

	current := dll.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	current.value = value

	return true
}

func (dll *DoublyLinkedList[T]) Remove(index int) bool {
	if index < 0 || index >= dll.size {
		return false
	}

	// Remove head
	if index == 0 {
		dll.head = dll.head.next
		if dll.head == nil {
			dll.tail = nil
		} else {
			dll.head.prev = nil
		}
		dll.size--
		return true
	}

	// Remove tail
	if index == dll.size-1 {
		dll.tail = dll.tail.prev
		if dll.tail == nil {
			dll.head = nil
		} else {
			dll.tail.next = nil
		}
		dll.size--
		return true
	}

	// Remove value in the middle
	current := dll.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	current.prev.next = current.next
	current.next.prev = current.prev
	dll.size--
	return true
}

func (dll *DoublyLinkedList[T]) Clear() {
	dll.head = nil
	dll.tail = nil
	dll.size = 0
}

func (dll *DoublyLinkedList[T]) Get(index int) (T, error) {
	var value T
	if index < 0 || index >= dll.size {
		return value, errors.New("index out of bound")
	}

	current := dll.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.value, nil
}

func (dll *DoublyLinkedList[T]) Size() int {
	return dll.size
}
