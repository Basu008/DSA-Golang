package main

import (
	"main/linkedlist"
)

func main() {
	var head *linkedlist.Node
	for i := 1; i <= 4; i++ {
		head = linkedlist.InsertNodeAtEnd(i, head)
	}
	linkedlist.DisplayDoublyLL(head)
	// fmt.Println("")
	head = linkedlist.DeleteLastNode(head)
	linkedlist.DisplayDoublyLL(head)
}
