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
