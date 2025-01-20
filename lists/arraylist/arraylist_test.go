package arraylist_test

import (
	"go-dsa/lists/arraylist"
	"testing"
)

func TestAdd(t *testing.T) {
	al := arraylist.New[int]()

	// 1. Empty
	if al.Size() != 0 {
		t.Errorf("Expected size 0, but got %d", al.Size())
	}

	// 2. Size 2
	al.Add(10)
	al.Add(20)
	if al.Size() != 2 {
		t.Errorf("Expected size 2, but got %d", al.Size())
	}
	val, ok := al.Get(0)
	if !ok {
		t.Errorf("Expected ok=true, but got false")
	}
	if val != 10 {
		t.Errorf("Expected 10, but got %d", val)
	}
	val, ok = al.Get(1)
	if !ok {
		t.Errorf("Expected ok=true, but got false")
	}
	if val != 20 {
		t.Errorf("Expected 20, but got %d", val)
	}
}

func TestRemove(t *testing.T) {
	al := arraylist.New[int]()

	// 1. Empty
	ok := al.Remove(100)
	if ok {
		t.Errorf("Expected ok=false, but ok=true")
	}

	// 2. Size 2
	al.Add(10)
	al.Add(20)
	ok = al.Remove(10)
	if !ok {
		t.Errorf("Expected ok=true, but got false")
	}
	if al.Size() != 1 {
		t.Errorf("Expected size 1, but got %d", al.Size())
	}
	val, ok := al.Get(0)
	if !ok {
		t.Errorf("expected ok=true, but got false")
	}
	if val != 20 {
		t.Errorf("Expected 20, but got %d", val)
	}
}
