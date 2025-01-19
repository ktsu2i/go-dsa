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
