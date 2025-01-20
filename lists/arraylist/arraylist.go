package arraylist

type ArrayList[T comparable] struct {
	values []T
}

func New[T comparable]() *ArrayList[T] {
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
