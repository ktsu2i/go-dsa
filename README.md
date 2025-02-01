# go-dsa (Go Data Structures and Algorithms)

This is an implementation of data structures and algorithms in Go.

## Purpose

The purpose of creating this module is as follows:

- Review data structures and algorithms
- Get familiar with Go
- Write test code
- Build a CI/CD pipeline

I started this project because I believe that creating a Go module like this is one of the good ways to achieve these goals.

## Installation

To install this module, just run these commands:

```
go get "github.com/ktsu2i/go-dsa"
go mod tidy
```

## Data Structures

- Lists
    - [Array List](#array-list)
    - [Singly Linked List](#singly-linked-list)
    - [Doubly Linked List](#doubly-linked-list)
- Queues
    - [Queue](#queue)
    - [Priority Queue](#priority-queue)
- [Set](#set)
- [Stack](#stack)
- Trees
    - [Binary Search Tree](#binary-search-tree)
    - [Min Heap](#min-heap)

### Array List

| Method | Parameters | Returns | Description |
| :---: | :---: | :---: | :---: |
| New | - | *ArrayList[T] | Creates and returns a new and empty array list for values of type `T`. |
| Add | value T | - | Adds the given value to the end of the array list. |
| Remove | index int | bool | Removes the value at the given index. Return `true` if successful, or `false` if the index is invalid. |
| Get | index int | (T, bool) | Gets the value at the given index. Returns the value and `true` if found; otherwise, returns the zero value and `false`. |
| Size | - | int | Returns the number of values in the array list. |
| Contains | value T | bool | Checks if the array list contains the given value; returns `true` if found, otherwise `false`. |
| MergeSort | - | - | Sorts the array list using merge sort |
| QuickSort | - | - | Sorts the array list using quick sort |

#### Usage

```go
package main

import "github.com/ktsu2i/go-dsa/lists/arraylist"

func main() {
    al := arraylist.New[int]()

    al.Add(3) // [3]
    al.Add(1) // [3, 1]
    al.Add(2) // [3, 1, 2]
    al.Add(4) // [3, 1, 2, 4]

    ok := al.Remove(3) // [3, 1, 2], ok: true
    
    val, ok := al.Get(0) // val: 3, ok: true
    
    size := al.Size() // size: 4

    exists := al.Contains(3) // exists: true

    al.MergeSort() // [1, 2, 3]
    al.QuickSort() // [1, 2, 3]
}
```

### Singly Linked List

| Method | Parameters | Returns | Description |
| :---: | :---: | :---: | :---: |
| New | - | *LinkedList[T] | Creates and returns a new and empty linked list for values of type `T`. |
| Add | value T | - | Adds the given value to the end of the linked list. |
| Set | index int, value T | bool | Sets the given value at the given index. Returns `true` if successful, otherwise `false` if the index is out of bounds. |
| Remove | index int | bool | Removes the value at the given index. Return `true` if successful, or `false` if the index is invalid. |
| Clear | - | - | Removes all elements from the linked list |
| Get | index int | (T, bool) | Gets the value at the given index. Returns the value and `true` if found; otherwise, returns the zero value and `false`. |
| Size | - | int | Returns the number of values in the linked list. |
| ToSlice | - | []T | Converts the linked list to a slice of type T, preserving the order of values. |

#### Usage

```go
package main

import "github.com/ktsu2i/go-dsa/lists/linkedlist"

func main() {
    ll := linkedlist.New[string]()

    ll.Add("A") // ["A"]
    ll.Add("B") // ["A", "B"]
    ll.Add("C") // ["A", "B", "C"]
    ll.Add("D") // ["A", "B", "C", "D"]

    ll.Set(3, "X") // ["A", "B", "C", "X"]
    
    ok := ll.Remove(3) // ["A", "B", "C"], ok: true
    
    val, ok := ll.Get(0) // val: "A", ok: true
    
    size := ll.Size() // size: 3

    slice := ll.ToSlice() // slice: [A B C]

    ll.Clear() // []
}
```

### Doubly Linked List

| Method | Parameters | Returns | Description |
| :---: | :---: | :---: | :---: |
| New | - | *DoublyLinkedList[T] | Creates and returns a new and empty doubly linked list for values of type `T`. |
| Add | value T | - | Adds the given value to the end of the doubly linked list. |
| Set | index int, value T | bool | Sets the given value at the given index. Returns `true` if successful, otherwise `false` if the index is out of bounds. |
| Remove | index int | bool | Removes the value at the given index. Return `true` if successful, or `false` if the index is invalid. |
| Clear | - | - | Removes all elements from the doubly linked list |
| Get | index int | (T, bool) | Gets the value at the given index. Returns the value and `true` if found; otherwise, returns the zero value and `false`. |
| Size | - | int | Returns the number of values in the doubly linked list. |
| ToSlice | - | []T | Converts the doubly linked list to a slice of type T, preserving the order of values. |

#### Usage

```go
package main

import "github.com/ktsu2i/go-dsa/lists/doublylinkedlist"

func main() {
    dll := doublylinkedlist.New[string]()

    dll.Add("A") // ["A"]
    dll.Add("B") // ["A", "B"]
    dll.Add("C") // ["A", "B", "C"]
    dll.Add("D") // ["A", "B", "C", "D"]

    dll.Set(3, "X") // ["A", "B", "C", "X"]
    
    ok := dll.Remove(3) // ["A", "B", "C"], ok: true
    
    val, ok := dll.Get(0) // val: "A", ok: true
    
    size := dll.Size() // size: 3

    slice := dll.ToSlice() // slice: [A B C]

    dll.Clear() // []
}
```

### Queue

| Method | Parameters | Returns | Description |
| :---: | :---: | :---: | :---: |
| New | - | *Queue[T] | Creates and returns a new and empty queue for values of type `T`. |
| Enqueue | value T | - | Adds the given value to the end of the queue. |
| Dequeue | - | (T, bool) | Removes and returns the front value of the queue. Returns `true` if successful, or `false` if empty. |
| Peek | - | (T, bool) | Returns the front value of the queue without removing it. Returns `true` if successful, or `false` if empty. |
| Size | - | int | Returns the number of values in the queue. |
| IsEmpty | - | bool | Returns `true` if empty, otherwise `false`. |

#### Usage

```go
package main

import "github.com/ktsu2i/go-dsa/queues/queue"

func main() {
    q := queue.New[string]()

    q.Enqueue("A") // ["A"]
    q.Enqueue("B") // ["A", "B"]
    q.Enqueue("C") // ["A", "B", "C"]
    q.Enqueue("D") // ["A", "B", "C", "D"]
    
    ok := q.Dequeue() // ["B", "C", "D"], ok: true
    
    front, ok := q.Peek() // front: "B", ok: true
    
    size := q.Size() // size: 3

    isEmpty := q.IsEmpty() // isEmpty: false
}
```

### Priority Queue

| Method | Parameters | Returns | Description |
| :---: | :---: | :---: | :---: |
| New | - | *PriorityQueue[T] | Creates and returns a new and empty priority queue for values of type `T`. |
| Enqueue | value T | - | Inserts the given value into the priority queue. |
| Dequeue | - | (T, bool) | Removes and returns the highest priority value from the priority queue. Returns `true` if successful, or `false` if empty. |
| Peek | - | (T, bool) | Returns the highest priority value without removing it. Returns `true` if successful, or `false` if the queue is empty. |
| Size | - | int | Returns the number of values in the priority queue. |
| IsEmpty | - | bool | Returns `true` if empty, otherwise `false`. |

#### Usage

```go
package main

import "github.com/ktsu2i/go-dsa/queues/priorityqueue"

func main() {
    pq := priorityqueue.New[string]()

    pq.Enqueue("A") // ["A"]
    pq.Enqueue("B") // ["A", "B"]
    pq.Enqueue("C") // ["A", "B", "C"]
    pq.Enqueue("D") // ["A", "B", "C", "D"]
    
    val, ok := pq.Dequeue() // val: "A", ok: true
    
    front, ok := pq.Peek() // front: "B", ok: true
    
    size := pq.Size() // size: 3

    isEmpty := pq.IsEmpty() // isEmpty: false
}
```

### Set

| Method | Parameters | Returns | Description |
| :---: | :---: | :---: | :---: |
| New | - | *Set[T] | Creates and returns a new and empty set for values of type `T`. |
| Add | value T | - | Adds the given value to the set. |
| Remove | value T | - | Removes the value from the set. |
| Contains | value T | bool | Checks if the set contains the given value; returns `true` if found, otherwise `false`. |
| Size | - | int | Returns the number of values in the set. |
| IsEmpty | - | bool | Returns `true` if empty, otherwise `false`. |
| Clear | - | - | Removes all elements from the set |
| Values | - | []T | Returns a slice of all values in the set |
| Union | other *Set[T] | *Set[T] | Returns the union of the two given sets. |
| Intersection | other *Set[T] | *Set[T] | Returns the intersection of the two given sets. |
| Difference | other *Set[T] | *Set[T] | Returns the difference of the two given sets. |
| Equals | other *Set[T] | *Set[T] | Returns `true` if the current set and the other set contain exactly the same values; otherwise, returns `false`. |

#### Usage

```go
package main

import "github.com/ktsu2i/go-dsa/set"

func main() {
    s1 := set.New[int]()

    s1.Add(1) // [1]
    s1.Add(2) // [1, 2]
    s1.Add(3) // [1, 2, 3]
    s1.Add(4) // [1, 2, 3, 4]
    
    val, ok := s1.Remove(3) // val: 4, ok: true
    
    exists := s1.Contains(1) // exists: 1
    
    size := s1.Size() // size: 3

    slice := s1.Values() // [1 2 3]

    s2 := set.New[int]()
    s2.Add(3)
    s2.Add(4)
    s2.Add(5)

    s := s1.Union(s2) // [1, 2, 3, 4, 5]
    s = s1.Intersection(s2) // [3]
    s = s1.Difference(s2) // [1, 2]

    isEqual := s1.Equals(s2) // isEqual: false

    s1.Clear() // []
    isEmpty := s1.IsEmpty() // isEmpty: true
}
```

### Stack

| Method | Parameters | Returns | Description |
| :---: | :---: | :---: | :---: |
| New | - | *Stack[T] | Creates and returns a new and empty stack for values of type `T`. |
| Push | value T | - | Adds the given value to the top of the stack. |
| Pop | - | (T, bool) | Removes and returns the top value of the queue. Returns `true` if successful, or `false` if empty. |
| Peek | - | (T, bool) | Returns the top value of the queue without removing it. Returns `true` if successful, or `false` if empty. |
| Size | - | int | Returns the number of values in the stack. |
| IsEmpty | - | bool | Returns `true` if empty, otherwise `false`. |

#### Usage

```go
package main

import "github.com/ktsu2i/go-dsa/stack"

func main() {
    s := stack.New[string]()

    s.Push("A") // ["A"]
    s.Push("B") // ["A", "B"]
    s.Push("C") // ["A", "B", "C"]
    s.Push("D") // ["A", "B", "C", "D"]
    
    val, ok := s.Pop() // val: "D", ok: true
    
    top, ok := s.Peek() // top: "C", ok: true
    
    size := s.Size() // size: 3

    isEmpty := s.IsEmpty() // isEmpty: false
}
```

### Binary Search Tree

| Method | Parameters | Returns | Description |
| :---: | :---: | :---: | :---: |
| New | - | *BinarySearchTree[T] | Creates and returns a new and empty binary search tree for values of type `T`. |
| Insert | value T | - | Inserts the given value into the binary search tree. |
| Contains | value T | bool | Checks if the binary search tree contains the given value. Returns `true` if found, otherwise `false`. |
| PreOrder | - | []T | Returns a slice containing the values in pre-order traversal order. |
| InOrder | - | int | Returns a slice containing the values in in-order traversal order. |
| PostOrder | - | bool | Returns a slice containing the values in post-order traversal order. |

#### Usage

```go
package main

import "github.com/ktsu2i/go-dsa/trees/bst"

func main() {
    bst := bst.New[string]()

    nums := []int{10, 5, 15, 2, 8, 12, 20}
	for _, v := range nums {
		bst.Insert(v)
	}
    
    exists := bst.Contains(10) // exists: true
    
    order := bst.PreOrder() // [10, 5, 2, 8, 15, 12, 20]
    order = bst.InOrder() // [2, 5, 8, 10, 12, 15, 20]
    order = bst.PostOrder() // [2, 8, 5, 12, 20, 15, 10]
}
```

### Min Heap

| Method | Parameters | Returns | Description |
| :---: | :---: | :---: | :---: |
| New | - | *MinHeap[T] | Creates and returns a new and empty min heap for values of type `T`. |
| Push | value T | - | Inserts the given value into the min heap. |
| Pop | - | (T, bool) | Removes and returns the root of the min heap. Returns `true` if successful, or `false` if empty. |
| Peek | - | (T, bool) | Returns the root of the min heap without removing it. Returns `true` if successful, or `false` if empty. |
| Size | - | int | Returns the number of values in the min heap. |
| IsEmpty | - | bool | Returns `true` if empty, otherwise `false`. |

#### Usage

```go
package main

import "github.com/ktsu2i/go-dsa/trees/minheap"

func main() {
    h := minheap.New[int]()

    nums := []int{10, 5, 15, 2, 8, 12, 20}
	for _, v := range nums {
		h.Push(v)
	}

    val, ok := h.Pop() // val: 2, ok: true
    
    val, ok = h.Peek() // val: 5, ok: true

    size := h.Size() // size: 6
    
    isEmpty := h.IsEmpty() // isEmpty: false
}
```
