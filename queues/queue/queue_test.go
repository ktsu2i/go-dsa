package queue_test

import (
	"go-dsa/queues/queue"
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

func TestDequeue(t *testing.T) {
	q := queue.New[string]()

	// 1. Dequeue from an empty queue
	_, ok := q.Dequeue()
	if ok {
		t.Errorf("Expected false after dequeueing from an empty queue, but got true")
	}
	if q.Size() != 0 {
		t.Errorf("Expected size 0, but got %d", q.Size())
	}

	q.Enqueue("A")
	q.Enqueue("B")

	// 2. Dequeue from queue with size 2
	val, ok := q.Dequeue()
	if !ok {
		t.Errorf("Expected true when dequeueing, but got false")
	}
	if val != "A" {
		t.Errorf("Expected A when dequeueing, but got %s", val)
	}
}

func TestPeek(t *testing.T) {
	q := queue.New[string]()

	// 1. Empty queue
	_, ok := q.Peek()
	if ok {
		t.Errorf("Expected false when peeking, but got true")
	}

	// 2. Size 1
	q.Enqueue("A")
	front, ok := q.Peek()
	if !ok {
		t.Errorf("Expected false when peeking, but got true")
	}
	if front != "A" {
		t.Errorf("Expected A on front, but got %s", front)
	}
	if q.Size() != 1 {
		t.Errorf("Expected size 1, but got %d", q.Size())
	}

	// 3. Size 2
	q.Enqueue("B")
	front, ok = q.Peek()
	if !ok {
		t.Errorf("Expected false when peeking, but got true")
	}
	if front != "A" {
		t.Errorf("Expected A on front, but got %s", front)
	}
	if q.Size() != 2 {
		t.Errorf("Expected size 2, but got %d", q.Size())
	}
}

func TestSize(t *testing.T) {
	q := queue.New[string]()

	// 1. Empty queue
	if q.Size() != 0 {
		t.Errorf("Expected size 0 for an empty queue, but got %d", q.Size())
	}

	// 2. Size 1
	q.Enqueue("A")
	if q.Size() != 1 {
		t.Errorf("Expected size 1 after enqueueing one value, but got %d", q.Size())
	}

	// 3. Reduce size -> Size 0
	q.Dequeue()
	if q.Size() != 0 {
		t.Errorf("Expected size 0, but got %d", q.Size())
	}
}

func TestIsEmpty(t *testing.T) {
	q := queue.New[string]()

	// 1. Initial empty queue
	if !q.IsEmpty() {
		t.Errorf("Expected size 0 for an empty queue, but got %d", q.Size())
	}

	// 2. Empty stack
	q.Enqueue("A")
	q.Dequeue()
	if !q.IsEmpty() {
		t.Errorf("Expected size 0 for an empty queue, but got %d", q.Size())
	}
}
