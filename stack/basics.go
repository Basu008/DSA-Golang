package stack

func AreBracketsValid(brackets string) bool {
	bracketsArr := []byte(brackets)
	stack := Stack{}
	for _, b := range bracketsArr {
		bracket := string(b)
		if isOpeningBracket(bracket) {
			stack.Push(bracket)
		} else {
			bracketFromStack := (stack.Pop()).(string)
			if bracketFromStack == "" || !areSameBracket(bracketFromStack, bracket) {
				return false
			}
		}
	}
	return true
}

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

func getNSL(arr []int) []int {
	stack := Stack{}
	res := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		top := stack.IntTop()
		for !stack.IsEmpty() && arr[top] >= arr[i] {
			stack.Pop()
			top = stack.IntTop()
		}
		if top == -1 {
			res[i] = i + 1
		} else {
			res[i] = i - top
		}
		stack.Push(i)
	}
	return res
}

func getNSR(arr []int) []int {
	stack := Stack{}
	res := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		top := stack.IntTop()
		for !stack.IsEmpty() && arr[top] >= arr[i] {
			stack.Pop()
			top = stack.IntTop()
		}
		if top == -1 {
			res[i] = len(arr) - i
		} else {
			res[i] = top - i
		}
		stack.Push(i)
	}
	return res
}

func isOpeningBracket(bracket string) bool {
	return bracket == "{" || bracket == "(" || bracket == "["
}

func areSameBracket(opening string, closing string) bool {
	return (opening == "{" && closing == "}") || (opening == "(" && closing == ")") || opening == "[" && closing == "]"
}
