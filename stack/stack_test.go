package stack_test

import (
	"go-dsa/stack"
	"testing"
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
