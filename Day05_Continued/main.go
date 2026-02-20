package main

import (
	"data-structure/queueslice"
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
}
