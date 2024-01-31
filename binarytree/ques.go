package binarytree

import (
	"fmt"

	"github.com/Basu008/DSA-Golang/queue"

	"math"

	"github.com/Basu008/DSA-Golang/stack"
)

type DataWithVertical struct {
	Node     *Node
	Position int
	Vertical int
}

type Pair struct {
	First  int
	Second int
}

func DepthOfTree(root *Node) int {
	if root == nil {
		return 0
	}
	leftHeight := DepthOfTree(root.Left)
	rightHeight := DepthOfTree(root.Right)
	return 1 + int(math.Max(float64(leftHeight), float64(rightHeight)))
}

func IsTreeBalanced(root *Node) bool {
	if root == nil {
		return true
	}
	leftHeight := DepthOfTree(root.Left)
	rightHeight := DepthOfTree(root.Right)
	var diff int
	if leftHeight >= rightHeight {
		diff = leftHeight - rightHeight
	} else {
		diff = rightHeight - leftHeight
	}
	if diff > 1 {
		return false
	}
	left := IsTreeBalanced(root.Left)
	right := IsTreeBalanced(root.Right)
	if !left || !right {
		return false
	}
	return true
}

func DiameterOfTree(root *Node, max *int) int {
	if root == nil {
		return 0
	}
	leftMax := DiameterOfTree(root.Left, max)
	rightMax := DiameterOfTree(root.Right, max)
	dia := leftMax + rightMax
	if *max <= dia {
		*max = dia
	}
	return 1 + int(math.Max(float64(leftMax), float64(rightMax)))
}

func MaxPathSum(root *Node, maxSum *int) int {
	if root == nil {
		return 0
	}
	leftSum := MaxPathSum(root.Left, maxSum)
	rightSum := MaxPathSum(root.Right, maxSum)
	sum := leftSum + rightSum + root.Data
	if *maxSum <= sum {
		*maxSum = sum
	}
	return root.Data + int(math.Max(float64(leftSum), float64(rightSum)))
}

func AreTreesEqual(root1 *Node, root2 *Node) bool {
	if root1 == nil || root2 == nil {
		return root1 == root2
	}
	return root1.Data == root2.Data && AreTreesEqual(root1.Left, root2.Left) &&
		AreTreesEqual(root1.Right, root2.Right)
}

func ZigZagTraversal(root *Node) {
	stack := stack.Stack{}
	queue := queue.Queue{}
	direction := "rtl"
	stack.Push(root)
	for !stack.IsEmpty() {
		for direction == "rtl" && !stack.IsEmpty() {
			topNode := stack.Pop().(*Node)
			if topNode.Right != nil {
				queue.Push(topNode.Right)
			}
			if topNode.Left != nil {
				queue.Push(topNode.Left)
			}
			fmt.Printf("%d ", topNode.Data)
		}
		direction = "ltr"
		for direction == "ltr" && !queue.IsEmpty() {
			topNode := queue.Pop().(*Node)
			if topNode.Right != nil {
				stack.Push(topNode.Right)
			}
			if topNode.Left != nil {
				stack.Push(topNode.Left)
			}
			fmt.Printf("%d ", topNode.Data)
		}
		direction = "rtl"
	}
}

func BoundaryTraversal(root *Node) {
	leftNodes, leafNodes, rightNodes := queue.Queue{}, queue.Queue{}, queue.Queue{}
	leftNodes.Push(root.Data)
	temp := root.Left
	for temp != nil {
		leftNodes.Push(temp.Data)
		if isPreeTreeNode(temp) {
			if temp.Left != nil {
				leafNodes.Push(temp.Left.Data)
			}
			if temp.Right != nil {
				leafNodes.Push(temp.Right.Data)
			}
			temp = nil
		}
	}
	temp = root.Right
	for temp != nil {
		rightNodes.Push(temp.Data)
		if isPreeTreeNode(temp) {
			if temp.Left != nil {
				leafNodes.Push(temp.Left.Data)
			}
			if temp.Right != nil {
				leafNodes.Push(temp.Right.Data)
			}
			temp = nil
		}
	}
	leftNodes.Display(true)
	leafNodes.Display(true)
	rightNodes.Display(false)
}

func TopView(root *Node) {
	queue := queue.Queue{}
	dataToVertical := make(map[int]int)
	initialNode := &DataWithVertical{
		Node:     root,
		Vertical: 0,
	}
	min, max := math.MaxInt, math.MinInt
	queue.Push(initialNode)
	for !queue.IsEmpty() {
		Top := queue.Pop().(*DataWithVertical)
		if _, ok := dataToVertical[Top.Vertical]; !ok {
			dataToVertical[Top.Vertical] = Top.Node.Data
			if min >= Top.Vertical {
				min = Top.Vertical
			}
			if max <= Top.Vertical {
				max = Top.Vertical
			}
		}
		if Top.Node.Left != nil {
			queue.Push(&DataWithVertical{
				Node:     Top.Node.Left,
				Vertical: Top.Vertical - 1,
			})
		}
		if Top.Node.Right != nil {
			queue.Push(&DataWithVertical{
				Node:     Top.Node.Right,
				Vertical: Top.Vertical + 1,
			})
		}
	}
	for i := min; i <= max; i++ {
		fmt.Printf("%d ", dataToVertical[i])
	}
}

func IsSymmetrical(root *Node) bool {
	return (root == nil) || isSymmetricalHelp(root.Left, root.Right)
}

func RootToNodePath(root *Node, res *[]int, key int) bool {
	if root == nil {
		return false
	}
	if root.Data == key {
		*res = append(*res, root.Data)
		return true
	}
	left := RootToNodePath(root.Left, res, key)
	if left {
		*res = append(*res, root.Data)
		return true
	}
	right := RootToNodePath(root.Right, res, key)
	if right {
		*res = append(*res, root.Data)
		return true
	}
	return false
}

func isPreeTreeNode(node *Node) bool {
	return (node.Left != nil && node.Left.Right == nil && node.Left.Left == nil) || (node.Right != nil &&
		node.Right.Left == nil && node.Right.Right == nil)
}

func isSymmetricalHelp(left *Node, right *Node) bool {
	if left == nil || right == nil {
		return left == right
	}
	if left.Data != right.Data {
		return false
	}
	return isSymmetricalHelp(left.Left, right.Right) && isSymmetricalHelp(left.Right, right.Left)
}

func LeastCommonAncestor(root *Node, lca Pair) *int {
	if root == nil {
		return nil
	}
	if root.Data == lca.First || root.Data == lca.Second {
		return &root.Data
	}
	left := LeastCommonAncestor(root.Left, lca)
	right := LeastCommonAncestor(root.Right, lca)
	if left != nil && right != nil {
		return &root.Data
	}
	if left != nil {
		return left
	}
	return right
}

func MaxWidthOfTree(root *Node) int {
	q := queue.Queue{}
	q.Push(DataWithVertical{
		Node:     root,
		Position: 0,
	})
	var maxWidth int
	for !q.IsEmpty() {
		size := q.Length()
		var first, last int
		min := q.Top().(DataWithVertical).Position
		for i := 0; i < size; i++ {
			currNode := q.Pop().(DataWithVertical)
			index := currNode.Position - min
			if i == 0 {
				first = currNode.Position
			} else if i == size-1 {
				last = currNode.Position
			}
			if currNode.Node.Left != nil {
				q.Push(DataWithVertical{
					Node:     currNode.Node.Left,
					Position: 2 * index,
				})
			}
			if currNode.Node.Right != nil {
				q.Push(DataWithVertical{
					Node:     currNode.Node.Right,
					Position: 2*index + 1,
				})
			}
		}
		maxWidth = int(math.Max(float64(maxWidth), float64(last-first+1)))
	}
	return maxWidth
}

func ChildrenSumProperty(root *Node) *Node {
	var sum int
	if root.Left == nil && root.Right == nil {
		return root
	}
	if root.Left != nil {
		sum += root.Left.Data
	}
	if root.Right != nil {
		sum += root.Right.Data
	}
	if root.Data > sum {
		if root.Left != nil {
			root.Left.Data = root.Data
		}
		if root.Right != nil {
			root.Right.Data = root.Data
		}
	} else {
		root.Data = sum
	}
	root.Left = ChildrenSumProperty(root.Left)
	root.Right = ChildrenSumProperty(root.Right)
	sum = 0
	if root.Left != nil {
		sum += root.Left.Data
	}
	if root.Right != nil {
		sum += root.Right.Data
	}
	root.Data = sum
	return root
}

func NodesAtKDistance(root *Node, key int, distance int) []int {
	res := []int{}
	var q, q2, visited queue.Queue
	q.Push(root)
	childToParentMap := make(map[*Node]*Node)
	for !q.IsEmpty() {
		topNode := q.Pop().(*Node)
		if topNode.Data == key {
			q2.Push(topNode)
			visited.Push(topNode)
		}
		if topNode.Left != nil {
			childToParentMap[topNode.Left] = topNode
			q.Push(topNode.Left)
		}
		if topNode.Right != nil {
			childToParentMap[topNode.Right] = topNode
			q.Push(topNode.Right)
		}
	}
	currentDistance := 0
	for !q2.IsEmpty() {
		size := q2.Length()
		if currentDistance == distance {
			for i := 0; i < size; i++ {
				node := q2.Pop().(*Node)
				res = append(res, node.Data)
			}
			return res
		}
		for i := 0; i < size; i++ {
			top := q2.Pop().(*Node)
			if parent, ok := childToParentMap[top]; ok {
				if !visited.Contains(parent) {
					q2.Push(parent)
					visited.Push(parent)
				}
			}
			if top.Left != nil {
				if !visited.Contains(top.Left) {
					q2.Push(top.Left)
					visited.Push(top.Left)
				}
			}
			if top.Right != nil {
				if !visited.Contains(top.Right) {
					q2.Push(top.Right)
					visited.Push(top.Right)
				}
			}
		}
		currentDistance++
	}
	return res
}

func TimeToBurnTree(root *Node, key int) int {
	var q queue.Queue
	var keyNode *Node
	var time int
	q.Push(root)
	childToParent := make(map[*Node]*Node)
	visitedNodes := make(map[*Node]int)
	for !q.IsEmpty() {
		topNode := q.Pop().(*Node)
		if topNode.Data == key {
			keyNode = topNode
		}
		if topNode.Left != nil {
			childToParent[topNode.Left] = topNode
			q.Push(topNode.Left)
		}
		if topNode.Right != nil {
			childToParent[topNode.Right] = topNode
			q.Push(topNode.Right)
		}
	}
	q.Push(keyNode)
	visitedNodes[keyNode] = 1
	for !q.IsEmpty() {
		size := q.Length()
		burned := false
		for i := 0; i < size; i++ {
			topNode := q.Pop().(*Node)
			if parentNode, ok := childToParent[topNode]; ok {
				if _, isVisited := visitedNodes[parentNode]; !isVisited {
					q.Push(parentNode)
					visitedNodes[parentNode] = 1
					burned = true
				}
			}
			if topNode.Left != nil {
				if _, isVisited := visitedNodes[topNode.Left]; !isVisited {
					q.Push(topNode.Left)
					visitedNodes[topNode.Left] = 1
					burned = true
				}
			}
			if topNode.Right != nil {
				if _, isVisited := visitedNodes[topNode.Right]; !isVisited {
					q.Push(topNode.Right)
					visitedNodes[topNode.Right] = 1
					burned = true
				}
			}
		}
		if burned {
			time++
		}
	}
	return time
}

func CountNodeForCompleteTree(root *Node) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	leftHeight := DepthOfTree(root.Left)
	rightHeight := DepthOfTree(root.Right)
	if leftHeight == rightHeight {
		return int(math.Pow(2, float64(leftHeight)) + 1)
	}
	leftNodeCount := CountNodeForCompleteTree(root.Left)
	rightNodeCount := CountNodeForCompleteTree(root.Right)
	return leftNodeCount + rightNodeCount + 1
}

func TreeFromInorderAndPreorder(preorder, inorder []int) *Node {
	pSize, iSize := len(preorder), len(inorder)
	if pSize == 0 || iSize == 0 {
		return nil
	}
	if iSize == 1 {
		return createNewNode(inorder[0])
	}
	node := createNewNode(preorder[0])
	leftInorder, rightInorder := []int{}, []int{}
	isLeft := true
	for _, iNode := range inorder {
		if iNode == preorder[0] {
			isLeft = false
			continue
		}
		if isLeft {
			leftInorder = append(leftInorder, iNode)
		} else {
			rightInorder = append(rightInorder, iNode)
		}
	}
	//{1, 2, 3, 4, 5, 6}
	node.Left = TreeFromInorderAndPreorder(preorder[1:1+len(leftInorder)], leftInorder)
	if len(preorder) > 1 {
		node.Right = TreeFromInorderAndPreorder(preorder[1+len(leftInorder):], rightInorder)
	}
	return node
}

func TreeFromInorderAndPostorder(postorder, inorder []int) *Node {
	pSize, iSize := len(postorder), len(inorder)
	if pSize == 0 || iSize == 0 {
		return nil
	}
	if iSize == 1 {
		return createNewNode(inorder[0])
	}
	node := createNewNode(postorder[pSize-1])
	leftInorder, rightInorder := []int{}, []int{}
	isLeft := true
	for _, iNode := range inorder {
		if iNode == postorder[pSize-1] {
			isLeft = false
			continue
		}
		if isLeft {
			leftInorder = append(leftInorder, iNode)
		} else {
			rightInorder = append(rightInorder, iNode)
		}
	}
	//{1, 2, 3, 4, 5, 6}
	node.Left = TreeFromInorderAndPostorder(postorder[:len(leftInorder)], leftInorder)
	if len(postorder) > 1 {
		node.Right = TreeFromInorderAndPostorder(postorder[len(leftInorder):pSize-1], rightInorder)
	}
	return node
}
