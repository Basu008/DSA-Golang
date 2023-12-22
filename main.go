package main

import (
	"fmt"
	"main/linkedlist"
)

func main() {
	var head *linkedlist.Node
	for i := 1; i <= 4; i++ {
		head = linkedlist.InsertNodeAtEnd(i, head)
	}
	linkedlist.DisplayDoublyLL(head)
	fmt.Println("")
	head = linkedlist.ReverseDLLRecursion(head)
	linkedlist.DisplayDoublyLL(head)
}
