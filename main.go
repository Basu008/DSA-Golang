package main

import (
	"github.com/Basu008/DSA-Golang/binarytree"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	var root *binarytree.Node
	for _, e := range arr {
		root = binarytree.Insert(root, e)
	}
	binarytree.DeSerializeTree("1,2,13,#,#,4,5,#,#,#,#").LevelOrder()
}
