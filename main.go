package main

import (
	"fmt"
	"main/linkedlist"
)

func main() {
	var head *linkedlist.Node
	arr := []int{9, 9, 9, 9, 9}
	for _, e := range arr {
		head = linkedlist.InsertNodeAtBeginning(e, head)
	}
	linkedlist.Display(head)
	fmt.Println("")
	head = linkedlist.IncrLL(head)
	linkedlist.Display(head)
}
