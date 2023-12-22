package linkedlist

import "fmt"

func InsertNodeAtEnd(data int, head *Node) *Node {
	d := NewNode(data)
	if head == nil {
		head = d
		return head
	}
	temp := head
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = d
	d.Prev = temp
	return head
}

func DeleteLastNode(head *Node) *Node {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return nil
	}
	temp := head
	for temp.Next.Next != nil {
		temp = temp.Next
	}
	temp.Next = nil
	return head
}

func DisplayDoublyLL(head *Node) {
	temp := head
	for temp != nil {
		fmt.Print(temp.Data, "<->")
		temp = temp.Next
	}
	fmt.Println("nil")
}

func ReverseDoublyLL(head *Node) *Node {
	if head == nil {
		return head
	}
	current := head
	for current.Next != nil {
		next := current.Next
		current.Next = current.Prev
		current.Prev = current.Next
		current = next
	}
	current.Next = current.Prev
	return current
}

func ReverseDLLRecursion(head *Node) *Node {
	if head == nil {
		return head
	}
	if head.Next == nil {
		head.Next = head.Prev
		return head
	}
	next := head.Next
	head.Next = head.Prev
	head.Prev = next
	return ReverseDLLRecursion(next)
}
