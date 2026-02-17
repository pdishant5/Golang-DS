package dynamicarray

import (
	"errors"
	"fmt"
)

type DynamicArray struct {
	data []int
	len  int
	cap  int
}

func New(initialCapacity int) (*DynamicArray, error) {
	if initialCapacity < 0 {
		return nil, errors.New("Invalid initial capacity! Provide a non-negative capacity.")
	}
	return &DynamicArray{
		data: make([]int, 0, initialCapacity),
		len:  0,
		cap:  initialCapacity,
	}, nil
}

func (da *DynamicArray) Append(value int) {
	if da.len == da.cap {
		newCap := 2 * da.cap
		if newCap == 0 {
			newCap = 1
		}

		newDA := make([]int, da.len+1, newCap)
		copy(newDA, da.data)

		newDA[da.len] = value
		da.data = newDA
		da.cap = newCap
	} else {
		da.data = da.data[:da.len+1]
		da.data[da.len] = value
	}
	da.len++
}

func (da *DynamicArray) Get(index int) int {
	if index < 0 || index >= da.len {
		panic(fmt.Sprintf("runtime error: index out of range [%d] with length %d", index, da.len))
	}
	return da.data[index]
}

func (da *DynamicArray) Set(index, value int) {
	if index < 0 || index >= da.len {
		panic(fmt.Sprintf("runtime error: index out of range [%d] with length %d", index, da.len))
	}

	da.data[index] = value
}

func (da DynamicArray) Data() []int {
	return da.data
}

func (da DynamicArray) Length() int {
	return da.len
}

func (da DynamicArray) Capacity() int {
	return da.cap
}
