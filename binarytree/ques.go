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
