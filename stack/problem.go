package stack

func StockSpan(stocks []int) []int {
	res := make([]int, len(stocks))
	stack := Stack{}
	for i := 0; i < len(stocks); i++ {
		curr := stocks[i]
		top := stack.IntTop()
		for top != -1 && stocks[top] <= curr {
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

func MaxAreaHistogram(histogram []int) int {
	var maxArea int
	leftResult := getNSL(histogram)
	rightResult := getNSR(histogram)
	for i, h := range histogram {
		left := leftResult[i]
		right := rightResult[i]
		w := (left + right - 1)
		area := h * w
		if maxArea <= area {
			maxArea = area
		}
	}
	return maxArea
}
