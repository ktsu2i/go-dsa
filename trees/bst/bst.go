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

func (bst *BinarySearchTree[T]) Insert(value T) {
	if bst.root == nil {
		bst.root = &node[T]{value: value}
	} else {
		insertNode(bst.root, value)
	}
}

func (bst *BinarySearchTree[T]) Contains(value T) bool {
	return containsNode(bst.root, value)
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

func (bst *BinarySearchTree[T]) PostOrder() []T {
	var result []T
	postOrderTraversal(bst.root, &result)
	return result
}

// Helper functions

func insertNode[T constraints.Ordered](n *node[T], value T) {
	if value < n.value {
		// left
		if n.left == nil {
			n.left = &node[T]{value: value}
		} else {
			insertNode(n.left, value)
		}
	} else {
		// right
		if n.right == nil {
			n.right = &node[T]{value: value}
		} else {
			insertNode(n.right, value)
		}
	}
}

func containsNode[T constraints.Ordered](n *node[T], value T) bool {
	if n == nil {
		return false
	}

	if value < n.value {
		return containsNode(n.left, value)
	} else if value > n.value {
		return containsNode(n.right, value)
	} else {
		return true
	}
}

func preOrderTraversal[T constraints.Ordered](n *node[T], result *[]T) {
	if n == nil {
		return
	}
	*result = append(*result, n.value)
	preOrderTraversal(n.left, result)
	preOrderTraversal(n.right, result)
}

func inOrderTraversal[T constraints.Ordered](n *node[T], result *[]T) {
	if n == nil {
		return
	}
	inOrderTraversal(n.left, result)
	*result = append(*result, n.value)
	inOrderTraversal(n.right, result)
}

func postOrderTraversal[T constraints.Ordered](n *node[T], result *[]T) {
	if n == nil {
		return
	}
	postOrderTraversal(n.left, result)
	postOrderTraversal(n.right, result)
	*result = append(*result, n.value)
}
