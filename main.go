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
	brackets := "()[{}()]"
	fmt.Print("Ans: ", stack.AreBracketsValid(brackets))
}
