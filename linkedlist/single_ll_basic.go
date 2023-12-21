package linkedlist

import (
	"fmt"
)

func InsertNode(data int, pos int, head *Node) *Node {
	if pos == 1 {
		return insertNodeAtBeginning(data, head)
	}
	return head
}

func DeleteNode(pos int, head *Node) *Node {
	length := Length(head)
	if pos == length {
		return deleteNodeAtLast(head)
	}
	return head
}

func IsExists(data int, head *Node) bool {
	temp := head
	for temp != nil {
		if temp.Data == data {
			return true
		}
		temp = temp.Next
	}
	return false
}

func Length(head *Node) int {
	size := 0
	temp := head
	for temp != nil {
		size++
		temp = temp.Next
	}
	return size
}

func Display(head *Node) {
	temp := head
	for temp != nil {
		fmt.Print(temp.Data, "->")
		temp = temp.Next
	}
	fmt.Print("nil")
}

func insertNodeAtBeginning(data int, head *Node) *Node {
	newNode := Node{
		Data: data,
	}
	if head == nil {
		head = &newNode
	} else {
		newNode.Next = head
		head = &newNode
	}
	return head
}

func deleteNodeAtLast(head *Node) *Node {
	temp := head
	if temp.Next == nil {
		head = nil
	}
	for temp.Next.Next != nil {
		temp = temp.Next
	}
	temp.Next = nil
	return head
}
