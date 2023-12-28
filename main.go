package main

import (
	"fmt"
	"main/linkedlist"
)

func main() {
	var head *linkedlist.Node
	arr := []int{10, 4, 10, 3, 5, 20, 10}
	for _, e := range arr {
		head = linkedlist.InsertNodeAtEnd(e, head)
	}
	fmt.Println("Initial List")
	linkedlist.DisplayDoublyLL(head)
	head = linkedlist.DeleteKeyNodes(head, 10)
	fmt.Println("Result")
	linkedlist.DisplayDoublyLL(head)
}
