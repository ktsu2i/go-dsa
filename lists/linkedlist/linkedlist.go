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
		current = ll.head.Next
	}

	current.Next = current.Next.Next

	if current.Next == nil {
		ll.tail = current
	}
	ll.size--
	return true
}
