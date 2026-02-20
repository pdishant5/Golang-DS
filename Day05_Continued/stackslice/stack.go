package stackslice

import "errors"

type Stack[T any] struct {
	Stack []T
	top   int
	size  int
	cap   int
}

func New[T any](cap int) *Stack[T] {
	return &Stack[T]{
		Stack: make([]T, cap),
		top:   -1,
		size:  0,
		cap:   cap,
	}
}

func (s *Stack[T]) Push(value T) error {
	if s.Full() {
		return errors.New("Stack is full!")
	}

	s.top++
	s.Stack[s.top] = value
	s.size++
	return nil
}

func (s *Stack[T]) Pop() error {
	if s.Empty() {
		return errors.New("Stack is empty!")
	}

	s.top--
	s.size--
	if s.Empty() {
		s.top = -1
	}
	return nil
}

func (s *Stack[T]) Top() (T, error) {
	var zero T
	if s.Empty() {
		return zero, errors.New("Stack is empty!")
	}
	return s.Stack[s.top], nil
}

func (s *Stack[T]) Empty() bool {
	return s.size == 0
}

func (s *Stack[T]) Full() bool {
	return s.size == s.cap
}
