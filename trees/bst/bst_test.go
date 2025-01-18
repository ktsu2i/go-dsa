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

func TestContains(t *testing.T) {
	bst := bst.New[int]()

	// 1. Empty BST
	if bst.Contains(10) {
		t.Errorf("Expected false checking if BST contains 10 on an empty BST, but got true")
	}

	values := []int{10, 5, 15, 2, 8, 12, 20}
	for _, v := range values {
		bst.Insert(v)
	}

	// 2. BST with 7 values
	for _, v := range values {
		if !bst.Contains(v) {
			t.Errorf("Expected to contain %d, but got true", v)
		}
	}

	// 3. Does NOT contain
	notInTree := []int{0, 6, 25, 100}
	for _, v := range notInTree {
		if bst.Contains(v) {
			t.Errorf("Expected not to contain %d, but got true", v)
		}
	}

	// 4. Duplicated value
	bst.Insert(10)
	if !bst.Contains(10) {
		t.Errorf("Expected to contain 10 after inserting the same value, but got false")
	}
}
