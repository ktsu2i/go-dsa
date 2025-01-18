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
