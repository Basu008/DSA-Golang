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
	arr := []int{3, 0, 2, 0, 4}
	fmt.Println("Rainwater trapping : ", stack.RainWaterTrapping(arr))
}
