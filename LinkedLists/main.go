package main

import (
	"fmt"
)

func main() {
	list := Create()

	for i := 0; i < 5; i++ {
		InsertAfter(list, i)
	}

	fmt.Println(*list)

	PrintList(list)
}
