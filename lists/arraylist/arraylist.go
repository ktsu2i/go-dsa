package arraylist

import "golang.org/x/exp/constraints"

type ArrayList[T constraints.Ordered] struct {
	values []T
}

func New[T constraints.Ordered]() *ArrayList[T] {
	return &ArrayList[T]{
		values: make([]T, 0),
	}
}

func (al *ArrayList[T]) Add(value T) {
	al.values = append(al.values, value)
}

func (al *ArrayList[T]) Remove(index int) bool {
	if index < 0 || index >= len(al.values) {
		return false
	}
	al.values = append(al.values[:index], al.values[index+1:]...)
	return true
}

func (al *ArrayList[T]) Get(index int) (T, bool) {
	var value T
	if index < 0 || index >= len(al.values) {
		return value, false
	}
	return al.values[index], true
}

func (al *ArrayList[T]) Size() int {
	return len(al.values)
}

func (al *ArrayList[T]) Contains(value T) bool {
	for _, v := range al.values {
		if v == value {
			return true
		}
	}
	return false
}

func (al *ArrayList[T]) MergeSort() {
	al.values = al.mergeSort(al.values)
}

func (al *ArrayList[T]) QuickSort() {
	al.quickSort(al.values, 0, len(al.values)-1)
}

// Helper functions

func (al *ArrayList[T]) mergeSort(arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := al.mergeSort(arr[:mid])
	right := al.mergeSort(arr[mid:])

	return al.merge(left, right)
}

func (al *ArrayList[T]) merge(left, right []T) []T {
	result := make([]T, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func (al *ArrayList[T]) quickSort(arr []T, low, high int) {
	if low < high {
		index := al.partition(arr, low, high)
		al.quickSort(arr, low, index-1)
		al.quickSort(arr, index+1, high)
	}
}

func (al *ArrayList[T]) partition(arr []T, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
