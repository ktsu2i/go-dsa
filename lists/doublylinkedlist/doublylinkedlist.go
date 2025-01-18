package doublylinkedlist

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
