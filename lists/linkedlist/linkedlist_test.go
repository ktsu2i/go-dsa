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
	ll.Add("D") // index: 3
}
