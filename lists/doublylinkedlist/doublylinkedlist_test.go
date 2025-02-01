package doublylinkedlist_test

import (
	"testing"

	"github.com/ktsu2i/go-dsa/lists/doublylinkedlist"
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

func TestRemove(t *testing.T) {
	dll := doublylinkedlist.New[string]()

	// 1. Remove from an empty list
	if dll.Remove(0) {
		t.Errorf("Expected false when removing from an empty list, but got true")
	}
	if dll.Size() != 0 {
		t.Errorf("Expected size 0, but got %d", dll.Size())
	}

	dll.Add("A") // index: 0
	dll.Add("B") // index: 1
	dll.Add("C") // index: 2
	dll.Add("D") // index: 3
	dll.Add("E") // index: 4

	// 2. Remove head
	ok := dll.Remove(0)
	if !ok {
		t.Errorf("Expected true when removing head, but got false")
	}
	// -> [B, C, D, E]
	if dll.Size() != 4 {
		t.Errorf("Expected size 4 after removing head, but got %d", dll.Size())
	}
	// Check if head is B at this point
	head, _ := dll.Get(0)
	if head != "B" {
		t.Errorf("Expected head to be B, but got %s", head)
	}

	// 3. Remove value at index 1
	ok = dll.Remove(1)
	if !ok {
		t.Errorf("Expected true when removing value at index 1, but got false")
	}
	// -> [B, D, E]
	if dll.Size() != 3 {
		t.Errorf("Expected size 3 after removing value at index 1, but got %d", dll.Size())
	}
	// Check if there are B at index 0, D at index 1, and E at index 2
	val1, _ := dll.Get(0)
	val2, _ := dll.Get(1)
	val3, _ := dll.Get(2)
	if val1 != "B" || val2 != "D" || val3 != "E" {
		t.Errorf("Expected [B, D, E], but got [%s, %s, %s]", val1, val2, val3)
	}

	// 4. Remove tail
	ok = dll.Remove(2)
	if !ok {
		t.Errorf("Expected true when removing head, but got false")
	}
	// -> [B, D]
	if dll.Size() != 2 {
		t.Errorf("Expected size 2 after removing tail, but got %d", dll.Size())
	}
	tail, _ := dll.Get(1)
	if tail != "D" {
		t.Errorf("Expected tail to be D, but got %s", tail)
	}

	// 5. Remove value at index out of bound
	ok = dll.Remove(9999)
	if ok {
		t.Errorf("Expected false for index out of bound, but got true")
	}
	// -> [B, D]
	if dll.Size() != 2 {
		t.Errorf("Expected size 2 after index out of bound error, but got %d", dll.Size())
	}
}

func TestClear(t *testing.T) {
	dll := doublylinkedlist.New[string]()

	// 1. Clear an empty list
	dll.Clear()
	if dll.Size() != 0 {
		t.Errorf("Expected size 0 after clearing an empty list, but got %d", dll.Size())
	}

	dll.Add("A")
	dll.Add("B")
	dll.Add("C")

	// 2. Clear a list with size 3
	dll.Clear()
	if dll.Size() != 0 {
		t.Errorf("Expected size 0 after clearing a list, but got %d", dll.Size())
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

func TestToSlice(t *testing.T) {
	dll := doublylinkedlist.New[string]()

	// 1. Empty list
	slice := dll.ToSlice()
	if len(slice) != 0 {
		t.Errorf("Expected an empty slice from an empty list, but got %v", slice)
	}

	dll.Add("A")
	dll.Add("B")
	dll.Add("C")

	// 2. [A, B, C]
	slice = dll.ToSlice()
	expected := []string{"A", "B", "C"}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %s at index %d, but got %s", v, i, slice[i])
		}
	}
}
