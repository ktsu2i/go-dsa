package queue

type Queue[T any] struct {
	data []T
}

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(value T) {
	q.data = append(q.data, value)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var value T
	if len(q.data) == 0 {
		return value, false
	}
	front := q.data[0]
	q.data = q.data[1:]
	return front, true
}

func (q *Queue[T]) Peek() (T, bool) {
	var value T
	if len(q.data) == 0 {
		return value, false
	}
	return q.data[0], true
}

func (q *Queue[T]) Size() int {
	return len(q.data)
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}
