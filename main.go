package main

import (
	"fmt"
	"main/linkedlist"
)

func main() {
	var head *linkedlist.Node
	arr := []int{4, 3, 2, 1}
	for _, e := range arr {
		head = linkedlist.InsertNodeAtBeginning(e, head)
	}
	fmt.Println("Initial List")
	linkedlist.Display(head)
	head = linkedlist.RotateLL(head, 5)
	fmt.Println("Result")
	linkedlist.Display(head)
}
