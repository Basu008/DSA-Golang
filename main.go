package main

import (
	"fmt"
	"main/linkedlist"
)

func main() {
	var head *linkedlist.Node
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	for _, e := range arr {
		head = linkedlist.InsertNodeAtBeginning(e, head)
	}
	fmt.Println("Initial List")
	linkedlist.Display(head)
	head = linkedlist.ReverseKSizeLL(head, 3)
	fmt.Println("Result")
	linkedlist.Display(head)
}
