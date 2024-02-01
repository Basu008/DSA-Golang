package binarytree

import (
	"fmt"
	"math"

	"github.com/Basu008/DSA-Golang/queue"
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

func GetNodesAtHeight(height int) int {
	return int(math.Pow(2, float64(height)) - 1)
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

func (root *Node) LevelOrder() {
	q := queue.Queue{}
	q.Push(root)
	fmt.Println("")
	for !q.IsEmpty() {
		top := q.Pop().(*Node)
		fmt.Printf("%d ", top.Data)
		if top.Left != nil {
			q.Push(top.Left)
		}
		if top.Right != nil {
			q.Push(top.Right)
		}
	}
}
