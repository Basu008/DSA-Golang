package main

import (
	"fmt"
	"main/linkedlist"
)

func main() {
	var head *linkedlist.Node
	arr := []int{1, 2, 2, 2}
	for _, e := range arr {
		head = linkedlist.InsertNodeAtEnd(e, head)
	}
	fmt.Println("Initial List")
	linkedlist.DisplayDoublyLL(head)
	head = linkedlist.RemoveDuplicates(head)
	fmt.Println("Result")
	linkedlist.DisplayDoublyLL(head)
}
