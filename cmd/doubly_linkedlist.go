package main

import (
	"fmt"
	"go-dsalgo/pkg/list"
)
func RunDoublyLinkedListImplementation() {

	fmt.Println("DoublyLinkedList Implementation")
	doublyLinkedList := list.NewDoublyLinkedList()
	doublyLinkedList.Append(1)
	doublyLinkedList.Append(2)
	doublyLinkedList.Unshift(3)


	fmt.Println(doublyLinkedList.Values())

}