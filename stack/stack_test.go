package stack_test

import (
	"testing"

	"github.com/ktsu2i/go-dsa/stack"
)

func TestPush(t *testing.T) {
	s := stack.New[string]()

	// 1. Push to an empty stack
	s.Push("A")
	if s.Size() != 1 {
		t.Errorf("Expected size 1 after first push, but got %d", s.Size())
	}
	top, ok := s.Peek()
	if !ok {
		t.Errorf("Expected true when peeking, but got false")
	}
	if top != "A" {
		t.Errorf("Expected A on top, but got %s", top)
	}

	// 2. Push to stack
	s.Push("B")
	if s.Size() != 2 {
		t.Errorf("Expected size 2 after pushing, but got %d", s.Size())
	}
	top, ok = s.Peek()
	if !ok {
		t.Errorf("Expected true when peeking, but got false")
	}
	if top != "B" {
		t.Errorf("Expected B on top, but got %s", top)
	}
}

func TestPop(t *testing.T) {
	s := stack.New[string]()

	// 1. Pop from an empty stack
	_, ok := s.Pop()
	if ok {
		t.Errorf("Expected false when popping from an empty stack, but got true")
	}
	if s.Size() != 0 {
		t.Errorf("Expected size 0, but got %d", s.Size())
	}

	s.Push("A")
	s.Push("B")
	s.Push("C")

	// 2. Pop from stack with size 3
	val, ok := s.Pop()
	if !ok {
		t.Errorf("Expected true when popping, but got false")
	}
	if val != "C" {
		t.Errorf("Expected C, but got %s", val)
	}
	if s.Size() != 2 {
		t.Errorf("Expected size 2, but got %d", s.Size())
	}

	// 3. Pop from stack with size 2
	val, ok = s.Pop()
	if !ok {
		t.Errorf("Expected true when popping, but got false")
	}
	if val != "B" {
		t.Errorf("Expected B, but got %s", val)
	}
	if s.Size() != 1 {
		t.Errorf("Expected size 1, but got %d", s.Size())
	}
}

func TestPeek(t *testing.T) {
	s := stack.New[string]()

	// 1. Empty stack
	_, ok := s.Peek()
	if ok {
		t.Errorf("Expected false when peeking, but got true")
	}

	// 2. Stack with size 1
	s.Push("A")
	top, ok := s.Peek()
	if !ok {
		t.Errorf("Expected false when peeking, but got true")
	}
	if top != "A" {
		t.Errorf("Expected A on top, but got %s", top)
	}
	if s.Size() != 1 {
		t.Errorf("Expected size 1, but got %d", s.Size())
	}

	// 3. Stack with size 2
	s.Push("B")
	top, ok = s.Peek()
	if !ok {
		t.Errorf("Expected false when peeking, but got true")
	}
	if top != "B" {
		t.Errorf("Expected B on top, but got %s", top)
	}
	if s.Size() != 2 {
		t.Errorf("Expected size 2, but got %d", s.Size())
	}
}

func TestSize(t *testing.T) {
	s := stack.New[string]()

	// 1. Empty stack
	if s.Size() != 0 {
		t.Errorf("Expected size 0 for an empty stack, but got %d", s.Size())
	}

	// 2. Size 1
	s.Push("A")
	if s.Size() != 1 {
		t.Errorf("Expected size 1 after pushing one value, but got %d", s.Size())
	}

	// 3. Reduce size -> Size 0
	s.Pop()
	if s.Size() != 0 {
		t.Errorf("Expected size 0, but got %d", s.Size())
	}
}

func TestIsEmpty(t *testing.T) {
	s := stack.New[string]()

	// 1. Initial empty stack
	if !s.IsEmpty() {
		t.Errorf("Expected size 0 for an empty stack, but got %d", s.Size())
	}

	// 2. Empty stack
	s.Push("A")
	s.Pop()
	if !s.IsEmpty() {
		t.Errorf("Expected size 0 for an empty stack, but got %d", s.Size())
	}
}
