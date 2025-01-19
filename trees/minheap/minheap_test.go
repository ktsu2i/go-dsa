package minheap_test

import (
	"go-dsa/trees/minheap"
	"testing"
)

func TestPush(t *testing.T) {
	h := minheap.New[int]()

	// 1. Empty heap
	if !h.IsEmpty() {
		t.Errorf("Expected an empty heap")
	}
	if h.Size() != 0 {
		t.Errorf("Expected size 0, but got %d", h.Size())
	}

	// 2. Heap with size 1
	h.Push(10)
	if h.Size() != 1 {
		t.Errorf("Expected size 1 after pushing one value, but got %d", h.Size())
	}
	val, ok := h.Peek()
	if !ok {
		t.Errorf("Expected true after peeking, but got false")
	}
	if val != 10 {
		t.Errorf("Expected 10 after peeking, but got %d", val)
	}

	h.Push(5)
	h.Push(15)

	// 3. Heap with size 3
	if h.Size() != 3 {
		t.Errorf("Expected size 3, but got %d", h.Size())
	}
	val, ok = h.Peek()
	if !ok {
		t.Errorf("Expected true after peeking, but got false")
	}
	if val != 5 {
		t.Errorf("Expected 5 after peeking, but got %d", val)
	}
}

func TestPop(t *testing.T) {
	h := minheap.New[int]()

	// 1. Empty heap
	_, ok := h.Pop()
	if ok {
		t.Errorf("Expected Pop() on empty heap to return ok=false, but got ok=true")
	}
	if h.Size() != 0 {
		t.Errorf("Expected size 0 for empty heap, but got %d", h.Size())
	}

	nums := []int{10, 5, 15, 2, 8, 12, 20}
	for _, v := range nums {
		h.Push(v)
	}
	if h.Size() != 7 {
		t.Errorf("Expected size 7, but got %d", h.Size())
	}

	// 2. Heap with size 7
	expected := []int{2, 5, 8, 10, 12, 15, 20}
	for _, exp := range expected {
		val, ok := h.Pop()
		if !ok {
			t.Errorf("Expected Pop() to return ok=true, but got ok=false")
			continue
		}
		if val != exp {
			t.Errorf("Expected Pop()=%d, but got %d", exp, val)
		}
	}
}

func TestPeek(t *testing.T) {
	h := minheap.New[int]()

	// 1. Empty heap
	_, ok := h.Peek()
	if ok {
		t.Errorf("Expected Peek() on empty heap to return ok=false, but got ok=true")
	}

	// 2. Heap with size 1
	h.Push(10)
	val, ok := h.Peek()
	if !ok {
		t.Errorf("Expected Peek() to return ok=false, but got ok=true")
	}
	if val != 10 {
		t.Errorf("Expected Peek()=10, but got %d", val)
	}
	if h.Size() != 1 {
		t.Errorf("Expected size 1, but got %d", h.Size())
	}

	// 3. Heap with size 3
	h.Push(5)
	h.Push(15)
	val, ok = h.Peek()
	if !ok {
		t.Errorf("Expected Peek() to return ok=false, but got ok=true")
	}
	if val != 5 {
		t.Errorf("Expected Peek()=5, but got %d", val)
	}
	if h.Size() != 3 {
		t.Errorf("Expected size 3, but got %d", h.Size())
	}
}

func TestSize(t *testing.T) {
	h := minheap.New[int]()

	// 1. Empty heap
	if h.Size() != 0 {
		t.Errorf("Expected size 0, but got %d", h.Size())
	}

	// 2. Heap with size 3
	h.Push(10)
	h.Push(5)
	h.Push(20)
	if h.Size() != 3 {
		t.Errorf("Expected size 3, but got %d", h.Size())
	}
}

func TestIsEmpty(t *testing.T) {
	h := minheap.New[int]()

	// 1. Initial empty heap
	if !h.IsEmpty() {
		t.Errorf("Expected size 0 for an empty heap, but got %d", h.Size())
	}

	// 2.Empty heap
	h.Push(5)
	h.Pop()
	if !h.IsEmpty() {
		t.Errorf("Expected size 0 for an empty heap, but got %d", h.Size())
	}
}
