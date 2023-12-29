package main

import (
	"fmt"
	"main/linkedlist"
)

func main() {
	var head *linkedlist.Node
	arr := []int{1, 0, 2, 1}
	for _, e := range arr {
		head = linkedlist.InsertNodeAtBeginning(e, head)
	}
	fmt.Println("Initial List")
	linkedlist.Display(head)
	head = linkedlist.Sort0s1s2s(head)
	fmt.Println("Result")
	linkedlist.Display(head)
}
