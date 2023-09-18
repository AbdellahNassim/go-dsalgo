package main

import (
	"fmt"
	"go-dsalgo/pkg/list"
)


func RunLinkedListImplementation() {
	fmt.Println("LinkedList Implementation")
	linkedList := list.NewLinkedList()
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)

	fmt.Println(linkedList.Values())
}