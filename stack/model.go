package stack

import "fmt"

type Stack []interface{}

func (s *Stack) Push(data interface{}) {
	*s = append(*s, data)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func (s *Stack) Display() {
	for i := len(*s) - 1; i >= 0; i-- {
		if i == len(*s)-1 {
			fmt.Print("[")
		}
		fmt.Print((*s)[i])
		if i == 0 {
			fmt.Print("]")
		} else {
			fmt.Print(",")
		}
	}
}

func (s *Stack) Top() interface{} {
	if s.IsEmpty() {
		return -1
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
