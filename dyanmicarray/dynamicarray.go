package dyanmicarray

import (
	"errors"
	"fmt"
)

type DyanmicArray[T comparable] struct {
	capacity int
	size     int
	arr      []T
}

func New[T comparable]() *DyanmicArray[T] {
	defaultCapacity := 8
	return &DyanmicArray[T]{capacity: defaultCapacity, arr: make([]T, defaultCapacity)}
}

func (da *DyanmicArray[T]) Append(val T) {
	if da.size < da.capacity {
		da.arr[da.size] = val
	} else {
		newCapacity := da.capacity * 2
		newArr := make([]T, newCapacity)
		copy(newArr, da.arr)
	}
	da.size += 1
}

func (da *DyanmicArray[T]) RemoveByVal(val T) (v T, err error) {
	matchedIndex := -1
	for i, v := range da.arr {
		if v == val {
			matchedIndex = i
			break
		}
	}
	if matchedIndex == -1 {
		return *new(T), errors.New("value not found")
	}
	for matchedIndex < da.size-1 {
		da.arr[matchedIndex] = da.arr[matchedIndex+1]
		matchedIndex += 1
	}
	return val, nil
}

func (da *DyanmicArray[T]) RemoveByIndex(index int) {
	if da.size <= index {
		return
	}
	for index < da.size-1 {
		da.arr[index] = da.arr[index+1]
		index += 1
	}
	da.size -= 1
}
func (da *DyanmicArray[T]) Print() {
	for i := 0; i < da.size; i++ {
		fmt.Println(da.arr[i])
	}
}
func (da *DyanmicArray[T]) IndexOf(val T) int {
	index := -1
	for i, v := range da.arr {
		if v == val {
			index = i
			break
		}
	}
	return index
}

func (da *DyanmicArray[T]) ValueAt(index int) T {
	return da.arr[index]
}

func (da *DyanmicArray[T]) Size() int {
	return da.size
}
