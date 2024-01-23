package main

import (
	"fmt"
	"main/binarytree"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	arr2 := []int{1, 2, 3, 4, 5}
	var root, root2 *binarytree.Node
	for _, e := range arr {
		root = binarytree.Insert(root, e)
	}
	for _, e := range arr2 {
		root2 = binarytree.Insert(root2, e)
	}
	fmt.Println("Are same", binarytree.AreTreesEqual(root, root2))
}
