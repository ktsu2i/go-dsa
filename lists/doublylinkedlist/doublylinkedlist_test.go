package doublylinkedlist_test

import (
	"go-dsa/lists/doublylinkedlist"
	"testing"
)

func TestAdd(t *testing.T) {
	dll := doublylinkedlist.New[string]()
	dll.Add("A")
	dll.Add("B")
	dll.Add("C")

	if dll.Size() != 3 {
		t.Errorf("Expected size 3, but got %d", dll.Size())
	}

	expected := []string{"A", "B", "C"}
	for i := 0; i < dll.Size(); i++ {
		if val, _ := dll.Get(i); val != expected[i] {
			t.Errorf("Expected %s at index %d,  but got %s", expected[i], i, val)
		}
	}
}

func TestGet(t *testing.T) {
	dll := doublylinkedlist.New[string]()

	// 1. Get from an empty list
	_, err := dll.Get(0)
	if err == nil {
		t.Errorf("Expected error when getting from an empty list, but got nil")
	}

	dll.Add("A")
	dll.Add("B")
	dll.Add("C")

	// 2. Get head
	head, err := dll.Get(0)
	if err != nil {
		t.Errorf("Unexpected error for Get(0): %v", err)
	}
	if head != "A" {
		t.Errorf("Expected head to be A, but got %s", head)
	}

	// 3. Get value at index 1
	val, err := dll.Get(1)
	if err != nil {
		t.Errorf("Unexpected error for Get(1): %v", err)
	}
	if val != "B" {
		t.Errorf("Expected B at index 1, but got %s", val)
	}

	// 4. Get tail
	tail, err := dll.Get(2)
	if err != nil {
		t.Errorf("Unexpected error for Get(2): %v", err)
	}
	if tail != "C" {
		t.Errorf("Expected tail to be C, but got %s", tail)
	}

	// 5. Get value at index out of bound
	_, err = dll.Get(9999)
	if err == nil {
		t.Errorf("Expected error for index out of bound, but got nil")
	}
}
