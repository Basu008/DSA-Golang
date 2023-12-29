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

func IncrLL(head *Node) *Node {
	if head == nil {
		return head
	}
	carry := getCarry(head)
	if carry > 0 {
		newNode := NewNode(carry)
		newNode.Next = head
		head = newNode
	}
	return head
}

func AddTwoLL(head1 *Node, head2 *Node) *Node {
	h1, h2 := head1, head2
	h3 := NewNode(-1)
	temp := h3
	carry := 0
	for h1 != nil || h2 != nil {
		sum := 0
		if h1 != nil {
			sum += h1.Data
		}
		if h2 != nil {
			sum += h2.Data
		}
		sum += carry
		carry = sum / 10
		newNode := NewNode(sum % 10)
		temp.Next = newNode
		temp = temp.Next
		if h1 != nil {
			h1 = h1.Next
		}
		if h2 != nil {
			h2 = h2.Next
		}
	}
	if carry == 1 {
		newNode := NewNode(carry)
		temp.Next = newNode
	}
	return h3.Next
}

func getCarry(node *Node) int {
	if node.Next == nil {
		sum := node.Data + 1
		carry := sum / 10
		node.Data = sum % 10
		return carry
	}
	carry := getCarry(node.Next)
	if carry > 0 {
		sum := node.Data + carry
		node.Data = sum % 10
		carry = sum / 10
	}
	return carry
}

func Sort0s1s2s(head *Node) *Node {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	zeroHead, oneHead, twoHead := NewNode(-1), NewNode(-1), NewNode(-1)
	zero, one, two := zeroHead, oneHead, twoHead
	temp := head
	for temp != nil {
		switch temp.Data {
		case 0:
			{
				zero.Next = temp
				zero = zero.Next
			}
		case 1:
			{
				one.Next = temp
				one = one.Next
			}
		case 2:
			{
				two.Next = temp
				two = two.Next
			}
		}
		temp = temp.Next
	}
	zero.Next = oneHead.Next
	one.Next = twoHead.Next
	two.Next = nil
	return zeroHead.Next
}
