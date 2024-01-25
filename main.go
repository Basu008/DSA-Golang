package main

import (
	"main/binarytree"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	var root *binarytree.Node
	for _, e := range arr {
		root = binarytree.Insert(root, e)
	}
	binarytree.ZigZagTraversal(root)
}
