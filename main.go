package main

import (
	"fmt"
	"main/linkedlist"
)

func main() {
	var head1, head2 *linkedlist.Node
	arr1 := []int{9, 9, 9, 9, 9, 9, 9}
	arr2 := []int{9, 9, 9, 9}
	for _, e := range arr1 {
		head1 = linkedlist.InsertNodeAtBeginning(e, head1)
	}
	for _, e := range arr2 {
		head2 = linkedlist.InsertNodeAtBeginning(e, head2)
	}
	fmt.Println("head 1")
	linkedlist.Display(head1)
	fmt.Println("head 2")
	linkedlist.Display(head2)
	fmt.Println("result")
	head := linkedlist.AddTwoLL(head1, head2)
	// head = linkedlist.IncrLL(head)
	linkedlist.Display(head)
}
