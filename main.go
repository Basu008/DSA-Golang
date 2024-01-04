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
	arr := []int{4, 5, 2, 10, 8}
	fmt.Println("NSR", stack.NSR(arr))
}
