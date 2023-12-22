package main

import (
	"fmt"
	"main/linkedlist"
)

func main() {
	var head *linkedlist.Node
	for i := 1; i <= 6; i++ {
		head = linkedlist.InsertNodeAtEnd(i, head)
	}
	linkedlist.DisplayDoublyLL(head)
	fmt.Println("")
	fmt.Println(linkedlist.MidOfLinkedList(head).Data)
}
