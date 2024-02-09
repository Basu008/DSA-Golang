package main

import (
	"github.com/Basu008/DSA-Golang/binarysearchtree"
)

func main() {
	arr := []int{7, 3, 15, 20, 9}
	var root *binarysearchtree.TreeNode
	for _, e := range arr {
		root = binarysearchtree.Insert(root, e)
	}
	root.Iterator([]string{"next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"})
}
