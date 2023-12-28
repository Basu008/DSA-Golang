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

func ReverseLLIterative(head *Node) *Node {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	curr := head
	var prev *Node
	next := curr.Next
	for curr.Next != nil {
		curr.Next = prev
		prev = curr
		curr = next
		next = curr.Next
	}
	curr.Next = prev
	return curr
}

func ReverseLLRecursion(head *Node) *Node {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	newHead := ReverseLLRecursion(head.Next)
	currNext := head.Next
	currNext.Next = head
	head.Next = nil
	return newHead
}

func DetectCycle(head *Node) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func StartOfCycle(head *Node) *Node {
	if head == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			start := head
			for start != slow {
				start = start.Next
				slow = slow.Next
			}
			return start
		}
	}
	return nil
}

func LengthOfLoop(head *Node) int {
	if head == nil {
		return 0
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			length := 1
			slow = slow.Next
			for slow != fast {
				slow = slow.Next
				length++
			}
			return length
		}
	}
	return 0
}

func OddEvenList(head *Node) *Node {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	odd := head
	even := head.Next
	evenHead := head.Next
	for even != nil && even.Next != nil {
		odd.Next = odd.Next.Next
		even.Next = even.Next.Next
		odd = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

func RemoveNthNode(head *Node, pos int) *Node {
	slow := head
	fast := head
	for i := 0; i < pos; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return head
}

func RemoveMiddleNode(head *Node) *Node {
	slow := head
	fast := head
	var slower *Node
	for fast != nil && fast.Next != nil {
		slower = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	slower.Next = slow.Next
	return head
}
