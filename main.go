package main

import (
	"main/stack"
)

func main() {
	st := stack.Stack{}
	for i := 5; i >= 0; i-- {
		st.Push(i)
	}
	st.Display()
}
