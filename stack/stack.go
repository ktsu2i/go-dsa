package stack

type Stack[T any] struct {
	data []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var value T
		return value, false
	}

	// Get the top value
	index := len(s.data) - 1
	value := s.data[index]
	s.data = s.data[:index]
	return value, true
}
