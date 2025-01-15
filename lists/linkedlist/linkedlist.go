package linkedlist

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (ll *LinkedList[T]) Add(value T) {
	node := &Node[T]{Value: value}

	if ll.head == nil {
		ll.head = node
		ll.tail = node
	} else {
		ll.head.Next = node
		ll.tail = node
	}

	ll.size++
}
