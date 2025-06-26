package linkedlist

import (
	"fmt"
)

type node[T comparable] struct {
	Val      T
	NextNode *node[T]
}
type LinkedList[T comparable] struct {
	size int
	head *node[T]
	tail *node[T]
}

func New[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{head: nil, tail: nil}
}

func (ll *LinkedList[T]) AddAtHead(val T) {
	newNode := node[T]{Val: val, NextNode: ll.head}
	ll.head = &newNode
	ll.size += 1
}
func (ll *LinkedList[T]) Add(val T) {
	newNode := node[T]{Val: val, NextNode: nil}
	if ll.size == 0 {
		ll.head = &newNode
		ll.tail = &newNode
	} else {
		ll.tail.NextNode = &newNode
		ll.tail = &newNode
	}
	ll.size += 1
}
func (ll *LinkedList[T]) ValueAtIndex(index int) T {
	currentNode := ll.head
	var val T
	for i := 0; currentNode != nil; i++ {
		if i == index {
			val = currentNode.Val
		}
		currentNode = currentNode.NextNode
	}
	return val
}

func (ll *LinkedList[T]) RemoveAtIndex(index int) T {
	var removedval T
	if ll.size <= index {
		return removedval
	}
	prevNode := ll.head
	currentNode := ll.head
	for i := 0; i <= index; i++ {
		if i == index {
			if prevNode == currentNode {
				removedval = currentNode.Val
				ll.head = currentNode.NextNode
			} else {
				removedval = currentNode.Val
				prevNode.NextNode = currentNode.NextNode
				currentNode.NextNode = nil
			}
			ll.size -= 1
			break
		}
		prevNode = currentNode
		currentNode = currentNode.NextNode
	}
	return removedval
}

func (ll *LinkedList[T]) Remove(val T) {
	currentNode := ll.head
	for i := 0; currentNode != nil; i++ {
		if currentNode.Val == val {
			ll.RemoveAtIndex(i)
			break
		}
		currentNode = currentNode.NextNode
	}
}

func (ll *LinkedList[T]) Size() int {
	return ll.size
}
func (ll *LinkedList[T]) Print() {
	currentNode := ll.head
	for currentNode != nil {
		fmt.Println(currentNode.Val)
		currentNode = currentNode.NextNode
	}
}
