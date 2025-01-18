package bst_test

import (
	"go-dsa/trees/bst"
	"testing"
)

func TestInsert(t *testing.T) {
	bst := bst.New[int]()

	// 1. Insert 1 value
	bst.Insert(10)
	inorder := bst.InOrder()
	expected := []int{10}
	if len(inorder) != len(expected) || inorder[0] != 10 {
		t.Errorf("Expected %v, but got %v", expected, inorder)
	}

	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(2)
	bst.Insert(8)
	bst.Insert(12)
	bst.Insert(20)

	// 2. Insert 7 values
	inorder = bst.InOrder()
	expected = []int{2, 5, 8, 10, 12, 15, 20}
	if len(inorder) != len(expected) {
		t.Errorf("Length mismatch, expected %v, but got %v", expected, inorder)
	}
	for i, v := range expected {
		if inorder[i] != v {
			t.Errorf("Expected %d at index %d, but got %d", v, i, inorder[i])
		}
	}

	// 3. Insert a duplicated number
	bst.Insert(10)
	inorder = bst.InOrder()
	expected = []int{2, 5, 8, 10, 10, 12, 15, 20}
	if len(inorder) != len(expected) {
		t.Errorf("Length mismatch, expected %v, but got %v", expected, inorder)
	}
	for i, v := range expected {
		if inorder[i] != v {
			t.Errorf("Expected %d at index %d, but got %d", v, i, inorder[i])
		}
	}
}
