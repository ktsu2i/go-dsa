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
