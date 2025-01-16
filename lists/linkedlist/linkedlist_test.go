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
			t.Errorf("At index %d, expected %s, got %s", i, expected[i], val)
		}
	}
}

func TestRemove(t *testing.T) {
	ll := linkedlist.New[string]()

	// 1. Remove from an empty list
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

	// 2. Remove head
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
		t.Errorf("Expected head to be B, but got %s", head)
	}

	// 3. Remove value at index 1
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
		t.Errorf("Expected [B, D, E], but got [%s, %s, %s]", val1, val2, val3)
	}

	// 4. Remove tail
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
		t.Errorf("Expected tail to be D, but got %s", tail)
	}

	// 5. Remove value at index out of bound
	ok = ll.Remove(9999)
	if ok {
		t.Errorf("Expected false for index out of bound, but got true")
	}
	// -> [B, D]
	if ll.Size() != 2 {
		t.Errorf("Expected size 2 after index out of bound error, but got %d", ll.Size())
	}
}

func TestGet(t *testing.T) {
	ll := linkedlist.New[string]()

	// 1. Get from an empty list
	_, err := ll.Get(0)
	if err == nil {
		t.Errorf("Expected error when getting from an empty list, but got nil")
	}

	ll.Add("A")
	ll.Add("B")
	ll.Add("C")

	// 2. Get head
	head, err := ll.Get(0)
	if err != nil {
		t.Errorf("Unexpected error for Get(0): %v", err)
	}
	if head != "A" {
		t.Errorf("Expected head to be A, but got %s", head)
	}

	// 3. Get value at index 1
	val, err := ll.Get(1)
	if err != nil {
		t.Errorf("Unexpected error for Get(1): %v", err)
	}
	if val != "B" {
		t.Errorf("Expected B at index 1, but got %s", val)
	}

	// 4. Get tail
	tail, err := ll.Get(2)
	if err != nil {
		t.Errorf("Unexpected error for Get(2): %v", err)
	}
	if tail != "C" {
		t.Errorf("Expected tail to be C, but got %s", tail)
	}

	// 5. Get value at index out of bound
	_, err = ll.Get(9999)
	if err == nil {
		t.Errorf("Expected error for index out of bound, but got nil")
	}
}

func TestSize(t *testing.T) {
	ll := linkedlist.New[string]()

	// 1. Empty list
	if ll.Size() != 0 {
		t.Errorf("Expected size 0 for an empty list, but got %d", ll.Size())
	}

	// 2. Size 1
	ll.Add("A") // [A]
	if ll.Size() != 1 {
		t.Errorf("Expected size 1 after adding one value, but got %d", ll.Size())
	}

	// 3. Size 3
	ll.Add("B") // [A, B]
	ll.Add("C") // [A, B, C]
	if ll.Size() != 3 {
		t.Errorf("Expected size 3 after adding three values, but got %d", ll.Size())
	}

	// 4. Reduce size -> Size 2
	ok := ll.Remove(1) // [A, C]
	if !ok {
		t.Errorf("Expected true when removing value at index 1, but got false")
	}
	if ll.Size() != 2 {
		t.Errorf("Expected size 2 after removing one value, but got %d", ll.Size())
	}
}
