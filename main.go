package main

import (
	"fmt"
	"main/binarytree"
)

func main() {
	// st := stack.Stack{}
	// for i := 5; i >= 0; i-- {
	// 	st.Push(i)
	// }
	// st.Display()
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	var root *binarytree.Node
	for _, e := range arr {
		root = binarytree.Insert(root, e)
	}
	binarytree.Inorder(root)
	fmt.Println("")
	binarytree.PreOrder(root)
	fmt.Println("")
	binarytree.PostOrder(root)
	binarytree.LevelOrder(root)
}
