# go-dsa (Go Data Structures and Algorithms)

This is an implementation of data structures and algorithms in Go.

## Purpose

The purpose of creating this module is as follows:

- Review data structures and algorithms
- Get familiar with Go
- Write test code
- Build a CI/CD pipeline

I started this project because I believe that creating a Go module like this is one of the good ways to achieve these goals.

## Data Structures

- Lists
    - [Array List](#array-list)
    - [Singly Linked List](#singly-linked-list)
    - [Doubly Linked List](#doubly-linked-list)
- Queues
    - [Queue]()
    - [Priority Queue]()
- [Set]()
- [Stack]()
- Trees
    - [Binary Search Tree]()
    - [Min Heap]()

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

### Queue

| Method | Parameters | Returns | Description |
| :---: | :---: | :---: | :---: |
| New | - | *Queue[T] | Creates and returns a new and empty queue for values of type `T`. |
| Enqueue | value T | - | Adds the given value to the end of the queue. |
| Dequeue | - | (T, bool) | Removes and returns the front value of the queue. Returns `true` if successful, or `false` if empty. |
| Peek | - | (T, bool) | Returns the front value of the queue without removing it. Returns `true` if successful, or `false` if empty. |
| Size | - | int | Returns the number of values in the queue. |
| IsEmpty | - | bool | Returns `true` if empty, otherwise `false`. |

### Priority Queue

| Method | Parameters | Returns | Description |
| :---: | :---: | :---: | :---: |
| New | - | *PriorityQueue[T] | Creates and returns a new and empty priority queue for values of type `T`. |
| Enqueue | value T | - | Inserts the given value into the priority queue. |
| Dequeue | - | (T, bool) | Removes and returns the highest priority value from the priority queue. Returns `true` if successful, or `false` if empty. |
| Peek | - | (T, bool) | Returns the highest priority value without removing it. Returns `true` if successful, or `false` if the queue is empty. |
| Size | - | int | Returns the number of values in the priority queue. |
| IsEmpty | - | bool | Returns `true` if empty, otherwise `false`. |

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

### Stack

| Method | Parameters | Returns | Description |
| :---: | :---: | :---: | :---: |
| New | - | *Stack[T] | Creates and returns a new and empty stack for values of type `T`. |
| Push | value T | - | Adds the given value to the top of the stack. |
| Pop | - | (T, bool) | Removes and returns the top value of the queue. Returns `true` if successful, or `false` if empty. |
| Peek | - | (T, bool) | Returns the top value of the queue without removing it. Returns `true` if successful, or `false` if empty. |
| Size | - | int | Returns the number of values in the stack. |
| IsEmpty | - | bool | Returns `true` if empty, otherwise `false`. |

### Binary Search Tree

| Method | Parameters | Returns | Description |
| :---: | :---: | :---: | :---: |
| New | - | *BinarySearchTree[T] | Creates and returns a new and empty binary search tree for values of type `T`. |
| Insert | value T | - | Inserts the given value into the binary search tree. |
| Contains | value T | bool | Checks if the binary search tree contains the given value. Returns `true` if found, otherwise `false`. |
| PreOrder | - | []T | Returns a slice containing the values in pre-order traversal order. |
| InOrder | - | int | Returns a slice containing the values in in-order traversal order. |
| PostOrder | - | bool | Returns a slice containing the values in post-order traversal order. |

### Min Heap

| Method | Parameters | Returns | Description |
| :---: | :---: | :---: | :---: |
| New | - | *MinHeap[T] | Creates and returns a new and empty min heap for values of type `T`. |
| Push | value T | - | Inserts the given value into the min heap. |
| Pop | - | (T, bool) | Removes and returns the root of the min heap. Returns `true` if successful, or `false` if empty. |
| Peek | - | (T, bool) | Returns the root of the min heap without removing it. Returns `true` if successful, or `false` if empty. |
| Size | - | int | Returns the number of values in the min heap. |
| IsEmpty | - | bool | Returns `true` if empty, otherwise `false`. |
