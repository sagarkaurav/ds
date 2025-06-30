package main

import (
	"fmt"

	"github.com/sagarkaurav/ds-golang/bst"
)

func main() {
	fmt.Println("data structures in Golang")
	bst := bst.New[int]()
	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(3)
	bst.Insert(0)
	bst.Delete(1)
	fmt.Println(bst.Contains(2))
	bst.LevelOrderPrint()
	// mh := minheap.BuildWithSlice([]int{7, 3, 1, 11, 0})
	// // mh.Print()
	// for !mh.IsEmpty() {
	// 	fmt.Println(mh.Poll())
	// }
	// queue := queue.New[int]()
	// queue.Enqueue(1)
	// queue.Enqueue(2)
	// queue.Enqueue(3)
	// queue.Dequeue()
	// queue.Print()
	// stack := stack.New[string]()
	// stack.Push("A")
	// stack.Push("B")
	// stack.Push("C")
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
