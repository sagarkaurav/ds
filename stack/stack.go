package stack

import (
	"github.com/sagarkaurav/ds/linkedlist"
)

type Stack[T comparable] struct {
	size int
	data *linkedlist.LinkedList[T]
}

func New[T comparable]() *Stack[T] {
	return &Stack[T]{data: linkedlist.New[T]()}
}
func (s *Stack[T]) Push(val T) {
	s.size += 1
	s.data.AddAtHead(val)
}

func (s *Stack[T]) Pop() T {
	s.size -= 1
	poppedVal := s.data.RemoveAtIndex(0)
	return poppedVal
}
func (s *Stack[T]) Peak() T {
	return s.data.ValueAtIndex(0)
}

func (s *Stack[T]) Print() {
	s.data.Print()
}

func (s *Stack[T]) Size() int {
	return s.size
}
