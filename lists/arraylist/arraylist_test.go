package arraylist_test

import (
	"fmt"
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
		t.Errorf("Expected ok=false, but true")
	}

	// 2. Size 2
	al.Add(10)
	al.Add(20)
	ok = al.Remove(0)
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

func TestGet(t *testing.T) {
	al := arraylist.New[int]()

	// 1. Empty
	_, ok := al.Get(0)
	if ok {
		t.Errorf("Expected ok=false, but got true")
	}

	// 2. Size 2
	al.Add(10)
	al.Add(20)
	val, ok := al.Get(0)
	if !ok {
		t.Errorf("expected ok=true, but got false")
	}
	if val != 10 {
		t.Errorf("Expected 10, but got %d", val)
	}
	val, ok = al.Get(1)
	if !ok {
		t.Errorf("expected ok=true, but got false")
	}
	if val != 20 {
		t.Errorf("Expected 20, but got %d", val)
	}
}

func TestContains(t *testing.T) {
	al := arraylist.New[int]()

	// 1. Empty
	if al.Contains(10) {
		t.Errorf("Expected to not contain 10, but it does")
	}

	// 2. Size 2
	al.Add(10)
	al.Add(20)
	if !al.Contains(10) {
		t.Errorf("Expected to contain 10, but it does not")
	}
	if !al.Contains(20) {
		t.Errorf("Expected to contain 20, but it does not")
	}
}

func TestMergeSort(t *testing.T) {
	al := arraylist.New[int]()

	// 1. Empty
	al.MergeSort()
	if al.Size() != 0 {
		t.Errorf("Expected size 0 after merge sort, but got %d", al.Size())
	}

	// 2. Size 1
	al.Add(5)
	al.MergeSort()
	val, _ := al.Get(0)
	if val != 5 {
		t.Errorf("Expected 5, but got %d", val)
	}

	// 3. Size 5
	al.Add(3)
	al.Add(1)
	al.Add(4)
	al.Add(2)
	fmt.Println(al)
	al.MergeSort()
	fmt.Println(al)
	expected := []int{1, 2, 3, 4, 5}
	for i, v := range expected {
		val, _ := al.Get(i)
		if val != v {
			t.Errorf("Expected %d at index %d, but got %d", v, i, val)
		}
	}
}
