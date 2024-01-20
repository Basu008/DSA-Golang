package queue

import "fmt"

type Queue []interface{}

func (q *Queue) Push(data interface{}) {
	*q = append(*q, data)
}

func (q *Queue) Pop() interface{} {
	if q.IsEmpty() {
		return nil
	}
	top := (*q)[0]
	*q = (*q)[1:]
	return top
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Display() {
	for i := len(*q) - 1; i >= 0; i-- {
		if i == len(*q)-1 {
			fmt.Print("[")
		}
		fmt.Print((*q)[i])
		if i == 0 {
			fmt.Print("]")
		} else {
			fmt.Print(",")
		}
	}
}

func (q *Queue) Top() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return (*q)[0]
}

func (q *Queue) IntTop() int {
	if q.IsEmpty() {
		return -1
	}
	return (*q)[0].(int)
}
