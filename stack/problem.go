package stack

import (
	"github.com/Basu008/DSA-Golang/arrays"
	"github.com/Basu008/DSA-Golang/basicmaths"
)

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

func MaxAreaInBinaryMatrix(binaryMatrix [][]int) int {
	histograms := make([][]int, len(binaryMatrix))
	for i := range binaryMatrix {
		currentHistogram := []int{}
		for j := range binaryMatrix[i] {
			if i == 0 {
				currentHistogram = append(currentHistogram, binaryMatrix[i][j])
			} else {
				if binaryMatrix[i][j] == 0 {
					currentHistogram = append(currentHistogram, 0)
				} else {
					currentHistogram = append(currentHistogram, histograms[i-1][j]+1)
				}
			}
		}
		histograms[i] = currentHistogram
	}
	maxAreas := []int{}
	for _, histogram := range histograms {
		maxAreas = append(maxAreas, MaxAreaHistogram(histogram))
	}
	return arrays.LargestElement(maxAreas)
}

func RainWaterTrapping(buildings []int) int {
	leftMaxArray := arrays.LargestElementLeft(buildings)
	rightMaxArray := arrays.LargestElementRight(buildings)
	var totalArea int
	for i, building := range buildings {
		area := basicmaths.Min(leftMaxArray[i], rightMaxArray[i]) - building
		totalArea += area
	}
	return totalArea
}
