package main

import (
	"fmt"
	"main/stack"
)

func main() {
	// st := stack.Stack{}
	// for i := 5; i >= 0; i-- {
	// 	st.Push(i)
	// }
	// st.Display()
	// arr := []int{0, 0, 0, 0, 0}
	binaryMatrix := [][]int{
		{1, 0, 1, 0, 0},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 1, 0}}
	fmt.Println("MAH in binary : ", stack.MaxAreaInBinaryMatrix(binaryMatrix))
}
