package linkedlist

type Node struct {
	Data int
	Next *Node
	Prev *Node
}

func NewNode(data int) *Node {
	return &Node{
		Data: data,
		Prev: nil,
		Next: nil,
	}
}
