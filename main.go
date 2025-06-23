package main

import (
	"fmt"

	"github.com/sagarkaurav/ds/dyanmicarray"
)

func main() {
	fmt.Println("data structures in Golang")
	da := dyanmicarray.New[int]()
	da.Append(1)
	da.Append(2)
	da.Append(3)
	da.Append(4)
	// da.RemoveByIndex(1)
	// da.RemoveByVal(1)
	// da.RemoveByVal(4)
	da.Print()
}
