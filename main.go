package main

import (
	"fmt"
	"main/binarytree"
)

func main() {
	arr := []int{3, 5, 1, 6, 2, 0, 8, 7, 4}
	var root *binarytree.Node
	for _, e := range arr {
		root = binarytree.Insert(root, e)
	}
	res := binarytree.NodesAtKDistance(root, 5, 2)
	fmt.Println(res)
}
