package stack

func NGR(arr []int) []int {
	stack := Stack{}
	res := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		for !stack.IsEmpty() && stack.Top().(int) <= arr[i] {
			stack.Pop()
		}
		var top int
		if stack.Top() == nil {
			top = -1
		} else {
			top = stack.Top().(int)
		}
		res[i] = top
		stack.Push(arr[i])
	}
	return res
}

func NGL(arr []int) []int {
	stack := Stack{}
	res := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		for !stack.IsEmpty() && stack.Top().(int) <= arr[i] {
			stack.Pop()
		}
		var top int
		if stack.Top() == nil {
			top = -1
		} else {
			top = stack.Top().(int)
		}
		res[i] = top
		stack.Push(arr[i])
	}
	return res
}

func NSL(arr []int) []int {
	stack := Stack{}
	res := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		for !stack.IsEmpty() && stack.Top().(int) >= arr[i] {
			stack.Pop()
		}
		var top int
		if stack.Top() == nil {
			top = -1
		} else {
			top = stack.Top().(int)
		}
		res[i] = top
		stack.Push(arr[i])
	}
	return res
}

func NSR(arr []int) []int {
	stack := Stack{}
	res := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		for !stack.IsEmpty() && stack.Top().(int) >= arr[i] {
			stack.Pop()
		}
		var top int
		if stack.Top() == nil {
			top = -1
		} else {
			top = stack.Top().(int)
		}
		res[i] = top
		stack.Push(arr[i])
	}
	return res
}
