package main

import (
	"fmt"

	"dynamic-array/dynamicarray"
)

func main() {
	da, err := dynamicarray.New(0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(da)
	for i := 1; i <= 10; i++ {
		da.Append(i)
		fmt.Printf("\nUnderlying array: %v\nLength: %d; Capacity: %d\n", da.Data(), da.Length(), da.Capacity())
	}
	// da.Set(11, 100) // panic: index out of bound
}
