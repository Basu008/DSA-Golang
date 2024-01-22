package binarytree

import "math"

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

func DiameterOfTree(root *Node, max []int) int {
	if root == nil {
		return 0
	}
	leftMax := DiameterOfTree(root.Left, max)
	rightMax := DiameterOfTree(root.Right, max)
	dia := leftMax + rightMax
	if max[0] <= dia {
		max[0] = dia
	}
	return 1 + int(math.Max(float64(leftMax), float64(rightMax)))
}
