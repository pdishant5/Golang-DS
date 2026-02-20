package queueslice

import "errors"

type Queue[T any] struct {
	Queue []T
	cap   int
	size  int
	front int
	rear  int
}

func New[T any](cap int) *Queue[T] {
	return &Queue[T]{
		Queue: make([]T, cap),
		cap:   cap,
		size:  0,
		front: -1,
		rear:  -1,
	}
}

func (q *Queue[T]) Push(value T) error {
	if q.Full() {
		return errors.New("Queue is full!")
	}

	if q.Empty() {
		q.front = 0
	}
	q.rear = (q.rear + 1) % q.cap
	q.Queue[q.rear] = value
	q.size++
	return nil
}

func (q *Queue[T]) Pop() error {
	if q.Empty() {
		return errors.New("Queue is empty!")
	}
	q.front = (q.front + 1) % q.cap
	q.size--

	if q.Empty() {
		q.front = -1
		q.rear = -1
	}
	return nil
}

func (q *Queue[T]) Front() (T, error) {
	var zero T
	if q.Empty() {
		return zero, errors.New("Queue is empty!")
	}
	return q.Queue[q.front], nil
}

func (q *Queue[T]) Empty() bool {
	return q.size == 0
}

func (q *Queue[T]) Full() bool {
	return q.size == q.cap
}
