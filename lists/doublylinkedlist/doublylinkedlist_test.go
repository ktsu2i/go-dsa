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

	expected := []string{"A", "B", "C"}
}
