package linkedlist

func ReverseKSizeLL(head *Node, k int) *Node {
	temp := head
	var next, prev *Node
	for temp != nil {
		kthNode := getKthNode(temp, k)
		if kthNode == nil {
			prev.Next = temp
			break
		}
		next = kthNode.Next
		kthNode.Next = nil
		kthNode = ReverseLLIterative(temp)
		if temp == head {
			head = kthNode
		} else {
			prev.Next = kthNode
		}
		prev = temp
		temp = next
	}
	return head
}

func getKthNode(head *Node, k int) *Node {
	temp := head
	for i := 1; i < k; i++ {
		temp = temp.Next
		if temp == nil {
			break
		}
	}
	return temp
}
