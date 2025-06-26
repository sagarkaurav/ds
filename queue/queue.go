package queue

import "github.com/sagarkaurav/ds/linkedlist"

type Queue[T comparable] struct {
	size int
	data *linkedlist.LinkedList[T]
}

func New[T comparable]() *Queue[T] {
	return &Queue[T]{data: linkedlist.New[T]()}
}

func (q *Queue[T]) Peak() T {
	return q.data.ValueAtIndex(0)
}

func (q *Queue[T]) Enqueue(val T) {
	q.size += 1
	q.data.Add(val)
}

func (q *Queue[T]) Dequeue() T {
	q.size -= 1
	return q.data.RemoveAtIndex(0)
}

func (q *Queue[T]) Size() int {
	return q.size
}

func (q *Queue[T]) Print() {
	q.data.Print()
}
