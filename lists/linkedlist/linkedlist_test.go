package linkedlist_test

import (
	"go-dsa/lists/linkedlist"
	"testing"
)

func TestAdd(t *testing.T) {
	ll := linkedlist.New[string]()
	ll.Add("A") // index: 0
	ll.Add("B") // index: 1
	ll.Add("C") // index: 2

	if ll.Size() != 3 {
		t.Errorf("Expected size 3, got %d", ll.Size())
	}

	expected := []string{"A", "B", "C"}
	for i := 0; i < ll.Size(); i++ {
		if val, _ := ll.Get(i); val != expected[i] {
			t.Errorf("At index %d, expected %v, got %v", i, expected[i], val)
		}
	}
}

func TestRemove(t *testing.T) {
	ll := linkedlist.New[string]()

	// Remove from an empty list
	if ll.Remove(0) {
		t.Errorf("Expected false when removing from an empty list, but got true")
	}
	if ll.Size() != 0 {
		t.Errorf("Expected size 0, but got %d", ll.Size())
	}

	ll.Add("A") // index: 0
	ll.Add("B") // index: 1
	ll.Add("C") // index: 2
	ll.Add("D") // index: 3
	ll.Add("E") // index: 4

	// Remove head
	ok := ll.Remove(0)
	if !ok {
		t.Errorf("Expected true when removing head, but got false")
	}
	// -> [B, C, D, E]
	if ll.Size() != 4 {
		t.Errorf("Expected size 4 after removing head, but got %d", ll.Size())
	}
	// Check if head is B at this point
	head, _ := ll.Get(0)
	if head != "B" {
		t.Errorf("Expected head to be B, but got %v", head)
	}

	// Remove value at index 1
	ok = ll.Remove(1)
	if !ok {
		t.Errorf("Expected true when removing value at index 1, but got false")
	}
	// -> [B, D, E]
	if ll.Size() != 3 {
		t.Errorf("Expected size 3 after removing value at index 1, but got %d", ll.Size())
	}
	// Check if there are B at index 0, D at index 1, and E at index 2
	val1, _ := ll.Get(0)
	val2, _ := ll.Get(1)
	val3, _ := ll.Get(2)
	if val1 != "B" || val2 != "D" || val3 != "E" {
		t.Errorf("Expected [B, D, E], but got [%v, %v, %v]", val1, val2, val3)
	}

	// Remove tail
	ok = ll.Remove(2)
	if !ok {
		t.Errorf("Expected true when removing head, but got false")
	}
	// -> [B, D]
	if ll.Size() != 2 {
		t.Errorf("Expected size 2 after removing tail, but got %d", ll.Size())
	}
	tail, _ := ll.Get(1)
	if tail != "D" {
		t.Errorf("Expected tail to be D, but got %v", tail)
	}

	// Remove value at index out of bound
	ok = ll.Remove(9999)
	if ok {
		t.Errorf("Expected false for index out of bound, but got true")
	}
	// -> [B, D]
	if ll.Size() != 2 {
		t.Errorf("Expected size 2 after index out of bound error, but got %d", ll.Size())
	}
}
