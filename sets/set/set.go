package set

type Set[T comparable] struct {
	values map[T]struct{}
}

func New[T comparable]() *Set[T] {
	return &Set[T]{values: make(map[T]struct{})}
}

func (s *Set[T]) Add(value T) {
	s.values[value] = struct{}{}
}

func (s *Set[T]) Remove(value T) {
	delete(s.values, value)
}

func (s *Set[T]) Contains(value T) bool {
	_, exists := s.values[value]
	return exists
}

func (s *Set[T]) Size() int {
	return len(s.values)
}

func (s *Set[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Set[T]) Clear() {
	s.values = make(map[T]struct{})
}

func (s *Set[T]) Values() []T {
	result := make([]T, 0, len(s.values))
	for v := range s.values {
		result = append(result, v)
	}
	return result
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	union := New[T]()
	for v := range s.values {
		union.Add(v)
	}
	for v := range other.values {
		union.Add(v)
	}
	return union
}

func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	intersection := New[T]()
	if s.Size() < other.Size() {
		for v := range s.values {
			if other.Contains(v) {
				intersection.Add(v)
			}
		}
	} else {
		for v := range other.values {
			if s.Contains(v) {
				intersection.Add(v)
			}
		}
	}
	return intersection
}

func (s *Set[T]) Equals(other *Set[T]) bool {
	if s.Size() != other.Size() {
		return false
	}
	for v := range s.values {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}
