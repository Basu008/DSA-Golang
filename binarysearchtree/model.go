package binarysearchtree

import (
	"fmt"
	"math"

	"github.com/Basu008/DSA-Golang/queue"
	"github.com/Basu008/DSA-Golang/stack"
)

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

func CreateNewNode(data int) *TreeNode {
	return &TreeNode{
		Data: data,
	}
}

func (tn *TreeNode) DoesExists(data int) bool {
	if tn == nil {
		return false
	}
	if data == tn.Data {
		return true
	}
	if data < tn.Data {
		return tn.Left.DoesExists(data)
	}
	return tn.Right.DoesExists(data)
}

func (tn *TreeNode) Ceil(data int, ceil []int) int {
	if tn == nil {
		return ceil[0]
	}
	if data == tn.Data {
		return data
	}
	if data < tn.Data {
		ceil[0] = tn.Data
		return tn.Left.Ceil(data, ceil)
	}
	return tn.Right.Ceil(data, ceil)
}

func (tn *TreeNode) Floor(data int, floor []int) int {
	if tn == nil {
		return floor[0]
	}
	if data == tn.Data {
		return data
	}
	if data < tn.Data {
		return tn.Left.Floor(data, floor)
	}
	floor[0] = tn.Data
	return tn.Right.Floor(data, floor)
}

func (tn *TreeNode) Delete(data int) *TreeNode {
	if tn == nil {
		return nil
	}
	if tn.Data == data {
		return helperFunction(tn)
	}
	temp := tn
	for temp != nil {
		if data < temp.Data {
			if temp.Left != nil && temp.Left.Data == data {
				temp.Left = helperFunction(temp.Left)
				break
			} else {
				temp = temp.Left
			}
		} else {
			if temp.Right != nil && temp.Right.Data == data {
				temp.Right = helperFunction(temp.Right)
				break
			} else {
				temp = temp.Right
			}
		}
	}
	return tn
}

func Insert(root *TreeNode, data int) *TreeNode {
	node := CreateNewNode(data)
	if root == nil {
		root = node
		return root
	}
	temp := root
	for temp != nil {
		if node.Data > temp.Data {
			if temp.Right == nil {
				temp.Right = node
				return root
			} else {
				temp = temp.Right
			}
		} else if node.Data < temp.Data {
			if temp.Left == nil {
				temp.Left = node
				return root
			} else {
				temp = temp.Left
			}
		}
	}
	return root
}

func (root *TreeNode) LevelOrder() {
	q := queue.Queue{}
	q.Push(root)
	fmt.Println("")
	for !q.IsEmpty() {
		top := q.Pop().(*TreeNode)
		fmt.Printf("%d ", top.Data)
		if top.Left != nil {
			q.Push(top.Left)
		}
		if top.Right != nil {
			q.Push(top.Right)
		}
	}
}

var count = 1

func (root *TreeNode) KthSmallestNode(pos int) int {
	if root == nil {
		return -1
	}
	var ans int
	ans = root.Left.KthSmallestNode(pos)
	if ans != -1 {
		return ans
	}
	if count == pos {
		return root.Data
	} else {
		count++
	}
	ans = root.Right.KthSmallestNode(pos)
	return ans
}

func (root *TreeNode) IsTreeBST() bool {
	return root.isTreeBST([]int{math.MinInt, math.MaxInt})
}

func (root *TreeNode) isTreeBST(limit []int) bool {
	if root == nil {
		return true
	}
	return (root.Data > limit[0] && root.Data < limit[1]) && root.Left.isTreeBST([]int{limit[0], root.Data}) && root.Right.isTreeBST([]int{root.Data, limit[1]})
}

func (root *TreeNode) LCAofBST(first, second int) int {
	if root == nil {
		return -1
	}
	if first < root.Data && second < root.Data {
		return root.Left.LCAofBST(first, second)
	}
	if first > root.Data && second > root.Data {
		return root.Right.LCAofBST(first, second)
	}
	return root.Data
}

func (root *TreeNode) Iterator(commands []string) {
	s := stack.Stack{}
	temp := root
	for temp != nil {
		s.Push(temp)
		temp = temp.Left
	}
	for _, command := range commands {
		if command == "next" {
			if s.IsEmpty() {
				fmt.Println("Error!")
				return
			}
			topNode := s.Pop().(*TreeNode)
			fmt.Printf("%v\n", topNode.Data)
			topNode = topNode.Right
			for topNode != nil {
				s.Push(topNode)
				topNode = topNode.Left
			}
		} else if command == "hasNext" {
			fmt.Printf("%v\n", !s.IsEmpty())
		}
	}
}

func (root *TreeNode) InorderSuccessor(data int) int {
	node := inorderHelper(root, data, nil)
	if node == nil {
		return -1
	}
	return node.Data
}

func inorderHelper(root *TreeNode, data int, head *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Data == data {
		if root.Right == nil {
			return head
		}
		return rightMinNode(root.Right)
	}
	if data < root.Data {
		return inorderHelper(root.Left, data, root)
	}
	return inorderHelper(root.Right, data, head)
}

func CreateTreeFromPreOrder(preorder []int) *TreeNode {
	root := createTree(preorder, []int{0}, math.MaxInt)
	return root
}

func createTree(preorder, pos []int, upperBound int) *TreeNode {
	if pos[0] == len(preorder) || preorder[pos[0]] > upperBound {
		return nil
	}
	root := CreateNewNode(preorder[pos[0]])
	pos[0] += 1
	root.Left = createTree(preorder, pos, root.Data)
	root.Right = createTree(preorder, pos, upperBound)
	return root
}

func helperFunction(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Left == nil {
		return node.Right
	}
	if node.Right == nil {
		return node.Left
	}
	leftMaxNode := leftMaxNode(node.Left)
	newRoot := node.Left
	leftMaxNode.Right = node.Right
	return newRoot
}

func leftMaxNode(node *TreeNode) *TreeNode {
	for node.Right != nil {
		node = node.Right
	}
	return node
}

func rightMinNode(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}
