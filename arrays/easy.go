package arrays

import (
	"fmt"
)

func LargestElement(arr []int) int {
	var max int
	for i := range arr {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}

func SecondLargestAndSmallest(arr []int) (int, int) {
	max := arr[0]
	min := arr[0]
	for i := range arr {
		if max < arr[i] {
			max = arr[i]
		}
		if min > arr[i] {
			min = arr[i]
		}
	}
	secondMin := arr[0]
	secondMax := arr[0]
	for i := range arr {
		if arr[i] > secondMax && arr[i] < max {
			secondMax = arr[i]
		}
		if secondMin > arr[i] && arr[i] > min {
			secondMin = arr[i]
		}
	}
	return secondMax, secondMin
}

func IsArraySorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func RemoveDuplicateElements(arr []int) int {
	size := len(arr)
	if size <= 1 {
		return size
	}
	k := 1
	curr := arr[0]
	for i := 0; i < len(arr)-1; i++ {
		if curr != arr[i] {
			curr = arr[i]
			arr[k], arr[i] = arr[i], arr[k]
			k++
			fmt.Println("Current: ", curr)
			fmt.Println("After swapping: ", arr)
		}
		// if arr[i] != arr[i+1] {
		// }
	}
	return k
}

func LeftRotateArray(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		arr[i], arr[i+1] = arr[i+1], arr[i]
	}
}

func RightRotateArray(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		arr[i], arr[i-1] = arr[i-1], arr[i]
	}
}

func RotateArrayNTimes(arr []int, k int, direction string) {
	if k == len(arr) {
		return
	}
	if direction == "left" {
		LeftRotateArrayNTimes(arr, k)
	} else {
		RightRotateArrayNTime(arr, k)
	}
}

func LeftRotateArrayNTimes(arr []int, k int) {
	size := len(arr)
	for i := k; i < size; i++ {
		arr[i], arr[i-k] = arr[i-k], arr[i]
	}
	if size%2 != 0 {
		RightRotateArray(arr[size-k : size])
	}
}

func RightRotateArrayNTime(arr []int, k int) {
	size := len(arr)
	for i := size - 1 - k; i >= 0; i-- {
		arr[i], arr[i+k] = arr[i+k], arr[i]
	}
	if size%2 != 0 {
		LeftRotateArray(arr[0:k])
	}
}

func MoveZerosToEnd(arr []int) {
	size := len(arr)
	if size <= 1 {
		return
	}
	i, j := 0, size-1
	for i < j {
		if arr[j] == 0 {
			j--
		} else if arr[i] != 0 {
			i++
		} else {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
}

func LinearSearch(arr []int, num int) int {
	for i := range arr {
		if arr[i] == num {
			return i
		}
	}
	return -1
}

func Union(arr1 []int, arr2 []int) []int {
	newArr := make([]int, 1)
	k := 0
	i, j := 0, 0
	len1, len2 := len(arr1), len(arr2)
	for i < len1 && j < len2 {
		if newArr[k] == arr1[i] {
			i++
		} else if newArr[k] == arr2[j] {
			j++
		} else if arr1[i] == arr2[j] {
			newArr = append(newArr, arr1[i])
			i++
			j++
			k++
		} else if arr1[i] < arr2[j] {
			newArr = append(newArr, arr1[i])
			i++
			k++
		} else {
			newArr = append(newArr, arr2[j])
			j++
			k++
		}
	}
	for i < len1 {
		if newArr[k] == arr1[i] {
			i++
		} else {
			newArr = append(newArr, arr1[i])
			i++
			k++
		}
	}
	for j < len2 {
		if newArr[k] == arr2[j] {
			j++
		} else {
			newArr = append(newArr, arr2[j])
			j++
			k++
		}
	}
	return newArr[1:]
}

func MissingNumberInArray(arr []int, N int) int {
	var sum int
	for _, element := range arr {
		sum += element
	}
	expectedSum := (N * (N + 1)) / 2
	return expectedSum - sum
}

func MaximumContiniousOnes(arr []int) int {
	var maxCount int
	var localCount int
	for _, value := range arr {
		if value == 1 {
			localCount++
		} else {
			localCount = 0
		}
		if localCount > maxCount {
			maxCount = localCount
		}
	}
	return maxCount
}

func FindSingleAppearanceElement(arr []int) int {
	hashMap := make(map[int]int)
	for _, element := range arr {
		hashMap[element]++
	}
	for key := range hashMap {
		if hashMap[key] == 1 {
			return key
		}
	}
	return -1
}

func LargestSubArrayWithGivenSum(arr []int, k int) int {
	size := len(arr)
	if size == 0 {
		return -1
	}
	start, end := 0, 0
	subArr := 0
	sum := arr[0]
	for end < size {
		for start <= end && sum > k {
			sum -= arr[start]
			start++
		}
		if sum == k {
			curr := end - start + 1
			if curr > subArr {
				subArr = curr
			}
		}
		end++
		if end < size {
			sum += arr[end]
		}
	}
	return subArr
}

func LongstSubArrayWithGivenSum2(arr []int, k int) int {
	size := len(arr)
	if size == 0 {
		return -1
	}
	hashMap := make(map[int]int)
	sum := 0
	maxLen := 0
	for i := range arr {
		sum += arr[i]
		if sum == k {
			curr := i + 1
			if curr > maxLen {
				maxLen = curr
			}
		}
		rem := sum - k
		value, exists := hashMap[rem]
		if exists {
			len := i - value
			if len > maxLen {
				maxLen = len
			}
		}
		_, empty := hashMap[sum]
		if !empty {
			hashMap[sum] = i
		}
	}
	return maxLen
}

func LargestElementLeft(arr []int) []int {
	res := make([]int, len(arr))
	for i, e := range arr {
		if i == 0 {
			res[i] = e
			continue
		}
		if res[i-1] <= e {
			res[i] = e
		} else {
			res[i] = res[i-1]
		}
	}
	return res
}

func LargestElementRight(arr []int) []int {
	res := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		if i == len(arr)-1 {
			res[i] = arr[i]
			continue
		}
		if res[i+1] <= arr[i] {
			res[i] = arr[i]
		} else {
			res[i] = res[i+1]
		}
	}
	return res
}
