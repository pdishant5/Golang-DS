package queuelinkedlist

import (
	"data-structure/node"
	"errors"
)

type Queue[T any] struct {
	Head *node.Node[T]
	Tail *node.Node[T]
	size int
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		Head: nil,
		Tail: nil,
		size: 0,
	}
}

func (q *Queue[T]) Push(value T) {
	node := node.New(value)

	if q.Empty() {
		q.Head = node
	} else {
		q.Tail.Next = node
	}
	q.Tail = node
	q.size++
}

func (q *Queue[T]) Pop() error {
	if q.Empty() {
		return errors.New("Queue is empty!")
	}
	node := q.Head
	q.Head = node.Next
	node.Next = nil
	q.size--

	if q.Empty() {
		q.Tail = nil
	}
	return nil
}

func (q *Queue[T]) Front() (T, error) {
	var zero T
	if q.Empty() {
		return zero, errors.New("Queue is empty!")
	}
	return q.Head.Data, nil
}

func (q *Queue[T]) Empty() bool {
	return q.size == 0
}
