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
