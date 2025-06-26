package main

import (
	"fmt"

	"github.com/sagarkaurav/ds/linkedlist"
)

func main() {
	fmt.Println("data structures in Golang")
	ll := linkedlist.New[int]()
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.Add(4)
	ll.Remove(1)
	ll.Print()
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
