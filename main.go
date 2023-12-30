package main

import (
	"fmt"
	"main/queue"
)

func main() {
	// st := stack.Stack{}
	// for i := 5; i >= 0; i-- {
	// 	st.Push(i)
	// }
	// st.Display()
	qu := queue.Queue{}
	for i := 5; i >= 0; i-- {
		qu.Push(i)
	}
	qu.Display()
	fmt.Printf("Removing from queue")
	qu.Pop()
	qu.Display()
}
