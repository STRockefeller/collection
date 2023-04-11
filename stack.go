package collection

type Stack[T any] struct {
	contents []T
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		contents: []T{},
	}
}

func NewFromSlice[T any](items []T) Stack[T] {
	if items == nil {
		return Stack[T]{
			contents: []T{},
		}
	}
	return Stack[T]{
		contents: items,
	}
}

func (s Stack[T]) IsEmpty() bool {
	return len(s.contents) == 0
}

// please make sure the stack is not empty or the method will raise a panic.
func (s Stack[T]) Peek() T {
	return s.contents[0]
}

// please make sure the stack is not empty or the method will raise a panic.
func (s *Stack[T]) Pop() T {
	res := s.Peek()
	s.contents = s.contents[1:]
	return res
}

func (s *Stack[T]) Push(item T) {
	s.contents = append([]T{item}, s.contents...)
}
