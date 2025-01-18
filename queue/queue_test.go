package queue_test

import (
	"go-dsa/queue"
	"testing"
)

func TestEnqueue(t *testing.T) {
	q := queue.New[string]()

	// 1. Enqueue one value
	q.Enqueue("A")
	if q.Size() != 1 {
		t.Errorf("Expected size 1, but got %d", q.Size())
	}
	front, ok := q.Peek()
	if !ok {
		t.Errorf("Expected true when peeking, but got false")
	}
	if front != "A" {
		t.Errorf("Expected front to be A, but got %s", front)
	}

	// 2. Enqueue two values
	q.Enqueue("B")
	if q.Size() != 2 {
		t.Errorf("Expected size 2, but got %d", q.Size())
	}
	front, ok = q.Peek()
	if !ok {
		t.Errorf("Expected true when peeking, but got false")
	}
	if front != "A" {
		t.Errorf("Expected front to be A, but got %s", front)
	}
}
