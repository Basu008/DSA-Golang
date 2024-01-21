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
