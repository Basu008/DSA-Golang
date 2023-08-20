package arrays

import (
	"fmt"
	"math"
)

func TwoSum(arr []int, k int) (string, int, int) {
	hashMap := make(map[int]int)
	for i := range arr {
		rem := k - arr[i]
		index, exists := hashMap[arr[i]]
		if exists {
			return "YES", index, i
		} else {
			hashMap[rem] = i
		}
	}
	return "NO", -1, -1
}

func Sort0s1s2s(arr []int) {
	size := len(arr)
	low, mid, high := 0, 0, size-1
	for mid <= high {
		if arr[mid] == 0 {
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
		}
		if arr[mid] == 1 {
			mid++
		}
		if arr[mid] == 2 {
			arr[mid], arr[high] = arr[high], arr[mid]
			high--
		}
	}
}

func MajorityElements(arr []int, size int) int {
	count, n := 0, arr[0]
	for _, value := range arr {
		if count == 0 {
			n = value
		}
		if n == value {
			count++
		} else {
			count--
		}
	}
	return n
}

func LargestSumSubArray(arr []int, size int) int {
	maxSum := math.MinInt64
	sum := 0
	for i := range arr {
		if sum < 0 {
			sum = 0
		}
		sum += arr[i]
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

func StockBuyAndSell(arr []int, size int) int {
	buyPrice := arr[0]
	maxProfit := 0
	for i := 1; i < size; i++ {
		profit := arr[i] - buyPrice
		if profit < 0 {
			buyPrice = arr[i]
		}
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}

func AlternatePositiveAndNegative(arr []int, size int) []int {
	var pos_arr = []int{}
	var neg_arr = []int{}
	var result = make([]int, size)
	for _, value := range arr {
		if value < 0 {
			neg_arr = append(neg_arr, value)
		} else {
			pos_arr = append(pos_arr, value)
		}
	}
	pos_size := len(pos_arr)
	neg_size := len(neg_arr)
	if pos_size > neg_size {
		for i := 0; i < neg_size; i++ {
			result[i*2] = pos_arr[i]
			result[i*2+1] = neg_arr[i]
		}
		index := neg_size * 2
		for i := neg_size; i < pos_size; i++ {
			result[index] = pos_arr[i]
			index += 1
		}
	} else {
		for i := 0; i < pos_size; i++ {
			result[i*2] = pos_arr[i]
			result[i*2+1] = neg_arr[i]
		}
		index := pos_size * 2
		for i := pos_size; i < neg_size; i++ {
			result[index] = pos_arr[i]
			index += 1
		}
	}

	return result
}

func FindNextPermutation(arr []int, size int) {
	pivot := -1
	for i := size - 2; i >= 0; i-- {
		if arr[i] < arr[i+1] {
			pivot = i
			break
		}
	}
	for i := size - 1; i > pivot; i-- {
		if arr[i] > arr[pivot] {
			arr[pivot], arr[i] = arr[i], arr[pivot]
			break
		}
	}
	ReverseArray(arr[pivot+1:])
}

func ReverseArray(arr []int) {
	size := len(arr)
	i, j := 0, size-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

func LeaderInArray(arr []int, size int) []int {
	if size <= 1 {
		return arr
	}
	result := []int{}
	max := arr[size-1]
	result = append(result, max)
	for i := size - 2; i >= 0; i-- {
		if arr[i] > max {
			result = append(result, arr[i])
			max = arr[i]
		}
	}
	return result
}

func SetZerosInMatrix(matrix [][]int, row int, column int) {
	column0 := 1
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				if j == 0 {
					column0 = 0
				} else {
					matrix[0][j] = 0
				}
			}
		}
	}
	for i := row - 1; i >= 0; i-- {
		for j := column - 1; j >= 0; j-- {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if column0 == 0 {
		matrix[0][0] = 0
	}
}

func RotateImageBy90(image [][]int) {
	for i := 0; i < len(image); i++ {
		for j := i; j < len(image[i]); j++ {
			image[i][j], image[j][i] = image[j][i], image[i][j]
		}
	}
	for i := 0; i < len(image); i++ {
		ReverseArray(image[i])
	}
}

func SpiralMatrixTransition(matrix [][]int) {
	right, down, left, up := 0, 0, len(matrix)-1, len(matrix[0])-1
	for down <= up {
		for i := right; i <= left; i++ {
			fmt.Print(matrix[down][i], " ")
		}
		down++
		for i := down; i <= up; i++ {
			fmt.Print(matrix[i][left], " ")
		}
		left--
		for i := left; i >= right; i-- {
			fmt.Print(matrix[up][i], " ")
		}
		up--
		for i := up; i >= down; i-- {
			fmt.Print(matrix[i][right], " ")
		}
		right++
	}
}

func CountSubArrayEqualToSum(arr []int, k int) int {
	count := 0
	sum := arr[0]
	size := len(arr)
	i, j := 0, 0
	for i < size && j < size {
		for i <= j && sum > k {
			sum -= arr[i]
			i++
		}
		if sum == k {
			count++
		}
		j++
		if j < size {
			sum += arr[j]
		}
	}
	return count
}
