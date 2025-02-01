package priorityqueue

import (
	"github.com/ktsu2i/go-dsa/trees/minheap"
	"golang.org/x/exp/constraints"
)

type PriorityQueue[T constraints.Ordered] struct {
	heap *minheap.MinHeap[T]
}

func New[T constraints.Ordered]() *PriorityQueue[T] {
	return &PriorityQueue[T]{heap: minheap.New[T]()}
}

func (pq *PriorityQueue[T]) Enqueue(value T) {
	pq.heap.Push(value)
}

func (pq *PriorityQueue[T]) Dequeue() (T, bool) {
	return pq.heap.Pop()
}

func (pq *PriorityQueue[T]) Peek() (T, bool) {
	return pq.heap.Peek()
}

func (pq *PriorityQueue[T]) Size() int {
	return pq.heap.Size()
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.heap.IsEmpty()
}
