package main

import (
	"fmt"
	"go-dsalgo/pkg/list"
)
func main() {
	fmt.Println("Hello World")

	l := list.NewSliceBasedList()
	l.Append(1)

	fmt.Println(l.Values())
}