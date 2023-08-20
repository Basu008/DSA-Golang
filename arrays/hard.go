package arrays

import (
	"fmt"
	"math"
	"sort"
)

func PascalTriangle(n int, r int, c int) {
	res := CalculateNCR(int64(r-1), int64(c-1))
	fmt.Println("Value at [", r, ",", c, "] is :", res)
	fmt.Println("Sequence at [", r, ",", c, "] is :", ShowNthRow(n))
	fmt.Println("Pattern at [", r, ",", c, "] is :")
	for i := 1; i <= n; i++ {
		fmt.Println(ShowNthRow(i))
	}
}

func ShowNthRow(row int) []int {
	product := 1
	res := []int{product}
	for i := 1; i < row; i++ {
		product *= (row - i)
		product /= i
		res = append(res, product)
	}
	return res
}

func CalculateNCR(n int64, r int64) int64 {
	var res int64 = 1
	var i int64 = 0
	for i < r {
		res *= (n - i)
		res /= (i + 1)
		i++
	}
	return res
}

func MajorityElementsN3(arr []int, size int) []int {
	count1, count2 := 0, 0
	element1, element2 := math.MinInt, math.MinInt
	for i := range arr {
		if count1 == 0 && element2 != arr[i] {
			element1 = arr[i]
			count1 = 1
		} else if count2 == 0 && element1 != arr[i] {
			element2 = arr[i]
			count2 = 1
		} else if arr[i] == element1 {
			count1++
		} else if arr[i] == element2 {
			count2++
		} else {
			count1--
			count2--
		}
	}
	count1, count2 = 0, 0
	min := size/3 + 1
	for i := range arr {
		if arr[i] == element1 {
			count1++
		} else if arr[i] == element2 {
			count2++
		}
	}
	result := []int{}
	if count1 >= min {
		result = append(result, element1)
	}
	if count2 >= min {
		result = append(result, element2)
	}
	return result
}

func ThreeSum(arr []int, size int) [][]int {
	result := [][]int{}
	if size < 3 {
		return result
	}
	sort.Ints(arr)
	for i := 0; i < size; i++ {
		if i != 0 && arr[i] == arr[i-1] {
			continue
		}
		j := i + 1
		k := size - 1
		for j < k {
			if arr[i]+arr[j]+arr[k] < 0 {
				j++
			} else if arr[i]+arr[j]+arr[k] > 0 {
				k--
			} else {
				indices := []int{arr[i], arr[j], arr[k]}
				result = append(result, indices)
				j++
				k--
				for j < k && arr[j-1] == arr[j] {
					j++
				}
				for j < k && arr[k+1] == arr[k] {
					k--
				}
			}
		}
	}
	return result
}

func FourSum(arr []int, size int, target int) [][]int {
	result := [][]int{}
	if size < 4 {
		return result
	}
	sort.Ints(arr)
	for i := 0; i < size; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		for j := i + 1; j < size; j++ {
			if j > i+1 && arr[j] == arr[j-1] {
				continue
			}
			k, l := j+1, size-1
			for k < l {
				sum := arr[i] + arr[j] + arr[k] + arr[l]
				if sum < target {
					k++
				} else if sum > target {
					l--
				} else {
					elements := []int{arr[i], arr[j], arr[k], arr[l]}
					result = append(result, elements)
					k++
					l--
					for k < l && arr[k] == arr[k-1] {
						k++
					}
					for k < l && arr[l] == arr[l+1] {
						l--
					}
				}
			}
		}
		fmt.Println("")
	}
	return result
}

func LongestSubArrayWithSumZero(arr []int) int {
	hashMap := make(map[int]int)
	maxLen := 0
	sum := 0
	for i := range arr {
		sum += arr[i]
		if sum == 0 {
			currentLen := i + 1
			if currentLen > maxLen {
				maxLen = currentLen
			}
		} else {
			value, exists := hashMap[sum]
			if exists {
				currentLen := i - value
				if currentLen > maxLen {
					maxLen = currentLen
				}
			} else {
				hashMap[sum] = i
			}
		}
	}
	return maxLen
}

func CalculateXOR(arr []int, k int) int {
	hashMap := make(map[int]int)
	count := 0
	xor := 0
	hashMap[xor] = 1
	for i := range arr {
		xor = xor ^ arr[i]
		xr := xor ^ k
		value, exists := hashMap[xr]
		if exists {
			count += value
		}
		v, ok := hashMap[xor]
		if ok {
			hashMap[xor] = v + 1
		} else {
			hashMap[xor] = 1
		}
	}
	return count
}

func MergeOverlappingIntervals(arr [][]int) [][]int {
	newArr := [][]int{}
	//initially we will sort the array
	newArr = append(newArr, arr[0])
	newArrSize := len(newArr)
	for i := 1; i < len(arr); i++ {
		if newArr[newArrSize-1][1] > arr[i][0] {
			if newArr[newArrSize-1][1] <= arr[i][1] {
				newArr[newArrSize-1] = []int{newArr[newArrSize-1][0], arr[i][1]}
			}
		} else {
			newArr = append(newArr, arr[i])
		}
		newArrSize = len(newArr)
	}
	return newArr
}

func MergeTwoSortedArrays(arr1 []int, size1 int, arr2 []int, size2 int) {
	i, j := size1-1, 0
	for i >= 0 && j < size2 {
		if arr1[i] >= arr2[j] {
			arr1[i], arr2[j] = arr2[j], arr1[i]
			i--
			j++
		} else {
			break
		}
	}
	sort.Ints(arr1)
	sort.Ints(arr2)
}

func FindMissingAndRepeatingNumber(arr []int, size int) (int, int) {
	sn := (size * (size + 1)) / 2
	s2n := (size * (size + 1) * (2*size + 1)) / 6
	s, s2 := 0, 0
	for _, value := range arr {
		s += value
		s2 += (value * value)
	}
	// x-y
	val1 := s - sn
	// x2-y2
	val2 := s2 - s2n
	// (x-y)(x+y) = s2 - s2n
	// (x+y)
	val2 = val2 / val1
	// X = ((x + y) + (x - y))/2
	// y = x - (x - y)
	x := (val2 + val1) / 2
	y := x - val1
	return y, x
}

func CountInversions(arr []int, start int, end int) int {
	count := 0
	if start >= end {
		return count
	}
	mid := (start + end) / 2
	count += CountInversions(arr, start, mid)
	count += CountInversions(arr, mid+1, end)
	count += Count(arr, start, mid, end)
	return count
}

func Count(arr []int, start int, mid int, end int) int {
	i, j := start, mid+1
	newArray := []int{}
	count := 0
	for i <= mid && j <= end {
		if arr[i] <= arr[j] {
			newArray = append(newArray, arr[i])
			i++
		} else {
			newArray = append(newArray, arr[j])
			count += (mid - i + 1)
			j++
		}
	}
	for i <= mid {
		newArray = append(newArray, arr[i])
		i++
	}
	for j <= end {
		newArray = append(newArray, arr[j])
		j++
	}
	for i := start; i <= end; i++ {
		arr[i] = newArray[i-start]
	}
	return count
}

func CountReversePairs(arr []int, start int, end int) int {
	count := 0
	if start >= end {
		return count
	}
	mid := (start + end) / 2
	count += CountReversePairs(arr, start, mid)
	count += CountReversePairs(arr, mid+1, end)
	count += CountReverse(arr, start, mid, end)
	Merge(arr, start, mid, end)
	return count
}

func CountReverse(arr []int, start int, mid int, end int) int {
	j := mid + 1
	count := 0
	for i := start; i <= mid; i++ {
		for j <= end && arr[i] > 2*arr[j] {
			j++
		}
		count += (j - (mid + 1))
	}
	return count
}

func Merge(arr []int, start int, mid int, end int) {
	i, j := start, mid+1
	newArray := []int{}
	for i <= mid && j <= end {
		if arr[i] <= arr[j] {
			newArray = append(newArray, arr[i])
			i++
		} else {
			newArray = append(newArray, arr[j])
			j++
		}
	}
	for i <= mid {
		newArray = append(newArray, arr[i])
		i++
	}
	for j <= end {
		newArray = append(newArray, arr[j])
		j++
	}
	for i := start; i <= end; i++ {
		arr[i] = newArray[i-start]
	}
}

func MaxProductSubArray(arr []int, n int) int {
	if len(arr) == 0 {
		return 0
	}
	maxProduct := math.MinInt
	pre, suff := 1, 1
	for i := range arr {
		if pre == 0 {
			pre = 1
		}
		if suff == 0 {
			suff = 1
		}
		pre *= arr[i]
		suff *= arr[n-i-1]
		if pre >= suff && pre >= maxProduct {
			maxProduct = pre
		} else if suff > pre && suff >= maxProduct {
			maxProduct = suff
		}
	}
	return maxProduct
}
