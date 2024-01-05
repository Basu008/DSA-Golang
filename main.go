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
	arr := []int{7, 2, 8, 9, 1, 3, 6, 5}
	fmt.Println("MAH : ", stack.MaxAreaHistogram(arr))
}
