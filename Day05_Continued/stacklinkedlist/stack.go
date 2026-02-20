package stacklinkedlist

import (
	"data-structure/node"
	"errors"
)

type Stack[T any] struct {
	top  *node.Node[T]
	size int
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		top:  nil,
		size: 0,
	}
}

func (s *Stack[T]) Push(value T) {
	node := node.New(value)

	if s.top != nil {
		node.Next = s.top
	}
	s.top = node
	s.size++
}

func (s *Stack[T]) Pop() error {
	if s.Empty() {
		return errors.New("Stack is empty!")
	}

	node := s.top
	s.top = node.Next
	node.Next = nil
	s.size--

	if s.Empty() {
		s.top = nil
	}
	return nil
}

func (s *Stack[T]) Top() (T, error) {
	var zero T
	if s.Empty() {
		return zero, errors.New("Stack is empty!")
	}
	return s.top.Data, nil
}

func (s *Stack[T]) Empty() bool {
	return s.size == 0
}
