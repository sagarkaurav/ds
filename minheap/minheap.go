package minheap

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type MinHeap[T constraints.Ordered] struct {
	data []T
}

func New[T constraints.Ordered]() *MinHeap[T] {
	return &MinHeap[T]{data: []T{}}
}

func BuildWithSlice[T constraints.Ordered](data []T) *MinHeap[T] {
	newHeap := &MinHeap[T]{}
	copy(newHeap.data, data)
	for mid := len(newHeap.data) / 2; mid >= 0; mid -= 1 {
		newHeap.sink(mid)
	}
	return newHeap
}

func (mh *MinHeap[T]) Insert(val T) {
	mh.data = append(mh.data, val)
	mh.swim(len(mh.data) - 1)
}

func (mh *MinHeap[T]) Poll() T {
	val := mh.data[0]
	mh.data[0] = mh.data[len(mh.data)-1]
	mh.data = mh.data[:len(mh.data)-1]
	mh.sink(0)
	return val
}

func (mh *MinHeap[T]) IsEmpty() bool {
	return len(mh.data) == 0
}

func (mh *MinHeap[T]) swim(currentIndex int) {
	parent_index := (currentIndex - 1) / 2
	for parent_index >= 0 && mh.data[parent_index] > mh.data[currentIndex] {
		mh.data[parent_index], mh.data[currentIndex] = mh.data[currentIndex], mh.data[parent_index]
		currentIndex = parent_index
		parent_index = (currentIndex - 1) / 2
	}
}
func (mh *MinHeap[T]) sink(currentIndex int) {
	smallest_val_index := currentIndex
	for {
		left := 2*currentIndex + 1
		right := 2*currentIndex + 2
		if left < len(mh.data) && mh.data[left] < mh.data[smallest_val_index] {
			smallest_val_index = left
		}
		if right < len(mh.data) && mh.data[right] < mh.data[smallest_val_index] {
			smallest_val_index = right
		}
		if smallest_val_index != currentIndex {
			mh.data[currentIndex], mh.data[smallest_val_index] = mh.data[smallest_val_index], mh.data[currentIndex]
			currentIndex = smallest_val_index
		} else {
			break
		}
	}
}

func (mh *MinHeap[T]) Print() {
	fmt.Println(mh.data)
}
