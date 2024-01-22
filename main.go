package main

import (
	"fmt"
	"main/binarytree"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var root *binarytree.Node
	for _, e := range arr {
		root = binarytree.Insert(root, e)
	}
	binarytree.LevelOrder(root)
	fmt.Println("Depth of tree", binarytree.DepthOfTree(root))
	max := []int{0}
	binarytree.DiameterOfTree(root, max)
	fmt.Println("Diameter", max)
}
