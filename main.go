package main

import (
	"fmt"
	"main/linkedlist"
)

func main() {
	var head *linkedlist.Node
	for i := 5; i > 0; i-- {
		head = linkedlist.InsertNodeAtBeginning(i, head)
	}
	linkedlist.Display(head)
	fmt.Println("")
	head = linkedlist.RemoveMiddleNode(head)
	linkedlist.Display(head)
}
