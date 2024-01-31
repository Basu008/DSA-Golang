package main

import (
	"main/binarytree"
)

func main() {
	// arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// var root *binarytree.Node
	// for _, e := range arr {
	// 	root = binarytree.Insert(root, e)
	// }
	root := binarytree.TreeFromInorderAndPostorder(
		[]int{40, 50, 20, 60, 30, 10}, []int{40, 20, 50, 10, 60, 30},
	)
	binarytree.LevelOrder(root)
}
