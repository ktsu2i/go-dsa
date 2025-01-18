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

func TestSet(t *testing.T) {
	dll := doublylinkedlist.New[string]()

	// 1. Set on an empty list
	ok := dll.Set(0, "A")
	if ok {
		t.Errorf("Expected false when setting a value on an empty list, but go true")
	}

	dll.Add("A")
	dll.Add("B")
	dll.Add("C") // -> [A, B, C]

	// 2. Set on head
	if !dll.Set(0, "X") {
		t.Errorf("Expected true when setting a value on head, but got false")
	}
	head, _ := dll.Get(0)
	if head != "X" {
		t.Errorf("Expected head to be X, but got %s", head)
	}

	// 3. Set at index 1
	if !dll.Set(1, "Y") {
		t.Errorf("Expected true when setting a value at index 1, but got false")
	}
	val, _ := dll.Get(1)
	if val != "Y" {
		t.Errorf("Expected Y at index 1, but got %s", val)
	}

	// 4. Set on tail
	if !dll.Set(2, "Z") {
		t.Errorf("Expected true when setting a value on tail, but got false")
	}
	tail, _ := dll.Get(2)
	if tail != "Z" {
		t.Errorf("Expected tail to be Z, but got %s", tail)
	}

	// 5. Set at index out of bound
	if dll.Set(9999, "Z") {
		t.Errorf("Expected false when setting at index out of bound, but got true")
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
