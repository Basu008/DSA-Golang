package binarytree

import (
	"fmt"
	"main/queue"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func createNewNode(data int) *Node {
	return &Node{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

func Insert(root *Node, data int) *Node {
	newNode := createNewNode(data)
	if root == nil {
		root = newNode
		return root
	}
	q := queue.Queue{}
	q.Push(root)
	for !q.IsEmpty() {
		node := q.Top().(*Node)
		if node.Left == nil {
			node.Left = newNode
			return root
		}
		if node.Right == nil {
			node.Right = newNode
			return root
		}
		q.Push(node.Left)
		q.Push(node.Right)
		q.Pop()
	}
	return root
}

func Inorder(root *Node) *Node {
	if root == nil {
		return root
	}
	Inorder(root.Left)
	fmt.Printf("%d ", root.Data)
	Inorder(root.Right)
	return root
}

func PreOrder(root *Node) *Node {
	if root == nil {
		return root
	}
	fmt.Printf("%d ", root.Data)
	PreOrder(root.Left)
	PreOrder(root.Right)
	return nil
}

func PostOrder(root *Node) *Node {
	if root == nil {
		return root
	}
	PostOrder(root.Left)
	PostOrder(root.Right)
	fmt.Printf("%d ", root.Data)
	return nil
}
