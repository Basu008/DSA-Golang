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

func (q *Queue) Display(reverse bool) {
	if reverse {
		for i := 0; i <= len(*q)-1; i++ {
			fmt.Print((*q)[i])
			fmt.Print(" ")
		}
	} else {
		for i := len(*q) - 1; i >= 0; i-- {
			fmt.Print((*q)[i])
			fmt.Print(" ")
		}
	}
}

func (q *Queue) Length() int {
	return len(*q)
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

func (q *Queue) Contains(e interface{}) bool {
	for _, v := range *q {
		if e == v {
			return true
		}
	}
	return false
}
