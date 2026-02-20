package main

import (
	"data-structure/queuelinkedlist"
	"data-structure/queueslice"
	"data-structure/stacklinkedlist"
	"data-structure/stackslice"
	"fmt"
)

func main() {
	// fixed-size circular queue
	q1 := queueslice.New[int](5)

	err := q1.Push(1)
	err = q1.Push(2)
	err = q1.Push(3)
	err = q1.Push(4)
	err = q1.Push(5)
	// no error till this point
	fmt.Println("Error:", err) // nil

	// accessing the front element
	val, err := q1.Front()
	if err != nil {
		fmt.Println("Error:", err) // nil
	}
	fmt.Println("Front element:", val) // 1

	// error - queue is full
	err = q1.Push(6)
	fmt.Println("Error:", err) // error

	err = q1.Pop()
	err = q1.Pop()
	// we can now push, because queue is not full
	err = q1.Push(6)
	fmt.Println("Error:", err) // nil

	val, err = q1.Front()
	if err != nil {
		fmt.Println("Error:", err) // nil
	}
	fmt.Println("Front element:", val) // 3

	err = q1.Pop()
	err = q1.Pop()
	err = q1.Pop()
	err = q1.Pop()
	// no error till this point - there is at least one value to pop
	fmt.Println("Error:", err) // nil

	// error - queue is empty
	err = q1.Pop()
	fmt.Println("Error:", err) // error

	val, err = q1.Front()
	if err != nil {
		fmt.Println("Error:", err) // error
	}
	fmt.Println("Front element:", val) // 0

	// infinite size queue
	q2 := queuelinkedlist.New[int]()

	q2.Push(1)
	q2.Push(2)
	q2.Push(3)
	q2.Push(4)
	q2.Push(5)
	// no error in pushing ever

	// no error in accessing the front element
	val, err = q2.Front()
	if err != nil {
		fmt.Println("Error:", err) // nil
	}
	fmt.Println("Front element:", val) // 1

	err = q2.Pop()
	err = q2.Pop()
	err = q2.Pop()

	val, err = q2.Front()
	if err != nil {
		fmt.Println("Error:", err) // nil
	}
	fmt.Println("Front element:", val) // 4

	err = q2.Pop()
	err = q2.Pop()
	// no error till this point
	fmt.Println("Error:", err)

	// err0r - queue is empty
	err = q2.Pop()
	fmt.Println("Error:", err)

	// error accessing front element, because queue is empty
	val, err = q2.Front()
	if err != nil {
		fmt.Println("Error:", err) // error
	}
	fmt.Println("Front element:", val) // 0

	// fixed-size stack
	s1 := stackslice.New[int](5)

	err = s1.Push(1)
	err = s1.Push(2)
	err = s1.Push(3)
	err = s1.Push(4)
	err = s1.Push(5)
	// no error till this point
	fmt.Println("Error:", err) // nil

	// accessing the top element
	val, err = s1.Top()
	if err != nil {
		fmt.Println("Error:", err) // nil
	}
	fmt.Println("Top element:", val) // 5

	// error - stack is full
	err = s1.Push(6)
	fmt.Println("Error:", err) // error

	err = s1.Pop()
	err = s1.Pop()
	// we can now push, because stack is not full
	err = s1.Push(6)
	fmt.Println("Error:", err) // nil

	val, err = s1.Top()
	if err != nil {
		fmt.Println("Error:", err) // nil
	}
	fmt.Println("Top element:", val) // 6

	err = s1.Pop()
	err = s1.Pop()
	err = s1.Pop()
	err = s1.Pop()
	// no error till this point - there is at least one value to pop
	fmt.Println("Error:", err) // nil

	// error - stack is empty
	err = s1.Pop()
	fmt.Println("Error:", err) // error

	val, err = s1.Top()
	if err != nil {
		fmt.Println("Error:", err) // error
	}
	fmt.Println("Top element:", val) // 0

	// infinite size stack
	s2 := stacklinkedlist.New[float64]()

	s2.Push(1)
	s2.Push(2)
	s2.Push(3)
	s2.Push(4)
	s2.Push(5)
	// no error till this point

	// accessing the top element - no error
	val2, err := s2.Top()
	if err != nil {
		fmt.Println("Error:", err) // nil
	}
	fmt.Println("Top element:", val2) // 5

	// no error
	s2.Push(6)

	err = s2.Pop()
	err = s2.Pop()
	// no error in pop
	fmt.Println("Error:", err) // nil

	val2, err = s2.Top()
	if err != nil {
		fmt.Println("Error:", err) // nil
	}
	fmt.Println("Top element:", val2) // 4

	err = s2.Pop()
	err = s2.Pop()
	err = s2.Pop()
	err = s2.Pop()
	// no error till this point - there is at least one value to pop
	fmt.Println("Error:", err) // nil

	// error - stack is empty
	err = s2.Pop()
	fmt.Println("Error:", err) // error

	val2, err = s2.Top()
	if err != nil {
		fmt.Println("Error:", err) // error
	}
	fmt.Println("Top element:", val2) // 0
}
