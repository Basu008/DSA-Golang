package linkedlist

func MidOfLinkedList(head *Node) *Node {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head.Next
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
