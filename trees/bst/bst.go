package bst

import "golang.org/x/exp/constraints"

type node[T constraints.Ordered] struct {
	value T
	left  *node[T]
	right *node[T]
}

type BinarySearchTree[T constraints.Ordered] struct {
	root *node[T]
}

func New[T constraints.Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{}
}

func (bst *BinarySearchTree[T]) Insert(v T) {
	if bst.root == nil {
		bst.root = &node[T]{value: v}
	} else {
		insertNode(bst.root, v)
	}
}

func (bst *BinarySearchTree[T]) Contains(v T) bool {
	return containsNode(bst.root, v)
}

func (bst *BinarySearchTree[T]) PreOrder() []T {
	var result []T
	preOrderTraversal(bst.root, &result)
	return result
}

func (bst *BinarySearchTree[T]) InOrder() []T {
	var result []T
	inOrderTraversal(bst.root, &result)
	return result
}

// Helper functions

func insertNode[T constraints.Ordered](n *node[T], v T) {
	if v < n.value {
		// left
		if n.left == nil {
			n.left = &node[T]{value: v}
		} else {
			insertNode(n.left, v)
		}
	} else {
		// right
		if n.right == nil {
			n.right = &node[T]{value: v}
		} else {
			insertNode(n.right, v)
		}
	}
}

func containsNode[T constraints.Ordered](n *node[T], v T) bool {
	if n == nil {
		return false
	}

	if v < n.value {
		return containsNode(n.left, v)
	} else if v > n.value {
		return containsNode(n.right, v)
	} else {
		return true
	}
}

func inOrderTraversal[T constraints.Ordered](n *node[T], result *[]T) {
	if n == nil {
		return
	}
	inOrderTraversal(n.left, result)
	*result = append(*result, n.value)
	inOrderTraversal(n.right, result)
}

func preOrderTraversal[T constraints.Ordered](n *node[T], result *[]T) {
	if n == nil {
		return
	}
	*result = append(*result, n.value)
	preOrderTraversal(n.left, result)
	preOrderTraversal(n.right, result)
}
