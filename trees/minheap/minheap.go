package minheap

import (
	"golang.org/x/exp/constraints"
)

type MinHeap[T constraints.Ordered] struct {
	data []T
}

func New[T constraints.Ordered]() *MinHeap[T] {
	return &MinHeap[T]{}
}

func (h *MinHeap[T]) Push(value T) {
	h.data = append(h.data, value)
	h.heapifyUp(len(h.data) - 1)
}

func (h *MinHeap[T]) Pop() (T, bool) {
	var value T
	if h.IsEmpty() {
		return value, false
	}

	min := h.data[0]
	last := len(h.data) - 1
	h.data[0] = h.data[last]
	h.data = h.data[:last]

	if !h.IsEmpty() {
		h.heapifyDown(0)
	}
	return min, true
}

func (h *MinHeap[T]) Peek() (T, bool) {
	var value T
	if h.IsEmpty() {
		return value, false
	}
	return h.data[0], true
}

func (h *MinHeap[T]) Size() int {
	return len(h.data)
}

func (h *MinHeap[T]) IsEmpty() bool {
	return len(h.data) == 0
}

// Helper functions

func (h *MinHeap[T]) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.data[index] < h.data[parentIndex] {
			h.data[index], h.data[parentIndex] = h.data[parentIndex], h.data[index]
			index = parentIndex
		} else {
			break
		}
	}
}

func (h *MinHeap[T]) heapifyDown(index int) {
	len := len(h.data)
	for {
		left := 2*index + 1
		right := 2*index + 2
		min := index

		if left < len && h.data[left] < h.data[min] {
			min = left
		}
		if right < len && h.data[right] < h.data[min] {
			min = right
		}
		if min != index {
			h.data[index], h.data[min] = h.data[min], h.data[index]
		} else {
			break
		}
	}
}
