package main

import (
	"fmt"

	"github.com/sagarkaurav/ds/minheap"
)

func main() {
	fmt.Println("data structures in Golang")
	mh := minheap.BuildWithSlice([]int{7, 3, 1, 11, 0})
	// mh.Print()
	for !mh.IsEmpty() {
		fmt.Println(mh.Poll())
	}
	// queue := queue.New[int]()
	// queue.Enqueue(1)
	// queue.Enqueue(2)
	// queue.Enqueue(3)
	// queue.Dequeue()
	// queue.Print()
	// stack := stack.New[int]()
	// stack.Push(1)
	// stack.Push(2)
	// stack.Push(3)
	// stack.Pop()
	// stack.Print()
	// ll := linkedlist.New[int]()
	// ll.Add(1)
	// ll.Add(2)
	// ll.Add(3)
	// ll.Add(4)
	// ll.AddAtHead(11)
	// ll.Print()
	// da := dyanmicarray.New[int]()
	// da.Append(1)
	// da.Append(2)
	// da.Append(3)
	// da.Append(4)
	// // da.RemoveByIndex(1)
	// // da.RemoveByVal(1)
	// // da.RemoveByVal(4)
	// da.Print()
}
