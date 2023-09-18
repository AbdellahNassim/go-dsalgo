package main

import (
	"fmt"
	"go-dsalgo/pkg/list"
)
func RunListImplementation() {
	fmt.Println("List implementation")
	slice := []interface{}{1, 2, 3, 4, 5}
	fmt.Println(slice)

	list := list.NewSliceBasedListFromSlice(slice);

	fmt.Println(list)

	list.Append(6)
	fmt.Println(list)

	list.Remove(0)
	fmt.Println(list)

	list.Unshift("Hello")
	fmt.Println(list)
}