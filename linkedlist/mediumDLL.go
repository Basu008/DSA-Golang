package linkedlist

func DeleteKeyNodes(head *Node, key int) *Node {
	temp := head
	for temp != nil {
		if temp.Data == key {
			if temp.Prev == nil {
				head = head.Next
			} else if temp.Next == nil {
				temp = temp.Prev
				temp.Next = nil
			} else {
				temp.Prev.Next = temp.Next
				temp.Next.Prev = temp.Prev
			}
		}
		temp = temp.Next
	}
	return head
}
