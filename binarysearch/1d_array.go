package binarysearch

import "math"

func FindElement(arr []int, start int, end int, k int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if arr[mid] == k {
		return mid
	} else if k < arr[mid] {
		return FindElement(arr, start, mid-1, k)
	} else {
		return FindElement(arr, mid+1, end, k)
	}
}

func FindLowerBound(arr []int, k int) int {
	start, end := 0, len(arr)-1
	ans := len(arr) - 1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] >= k {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return ans
}

func FindUpperBound(arr []int, k int) int {
	start, end := 0, len(arr)-1
	ans := len(arr)
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] > k {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return ans
}

func SearchInsertPosition(arr []int, k int) int {
	start, end := 0, len(arr)-1
	ans := len(arr) - 1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == k {
			return mid
		} else if arr[mid] > k {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return ans
}

func FloorAndCeilInArray(arr []int, k int) (int, int) {
	start, end := 0, len(arr)-1
	ans := len(arr) - 1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == k {
			return arr[mid], arr[mid]
		} else if arr[mid] > k {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	floor, ceil := -1, ans
	if ans > 0 {
		floor = ans - 1
	}
	return arr[floor], arr[ceil]
}

func FindLastIndex(arr []int, k int) int {
	start, end := 0, len(arr)-1
	ans := -1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] <= k {
			if arr[mid] == k {
				ans = mid
			}
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return ans
}

func CountOccurance(arr []int, k int) int {
	firstOccurance := FindLowerBound(arr, k)
	lastOccurance := FindUpperBound(arr, k)
	return lastOccurance - firstOccurance
}

func FindInRotatedArray(arr []int, k int) int {
	start, end := 0, len(arr)-1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == k {
			return mid
		} else if arr[mid] > k {
			if arr[start] <= arr[mid] && isInRange(arr[start], k, arr[mid]) {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if arr[mid] <= arr[end] && isInRange(arr[mid], k, arr[end]) {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}

func isInRange(a int, b int, c int) bool {
	if b >= a && b <= c {
		return true
	}
	return false
}

func FindInRotatedArray2(arr []int, k int) bool {
	start, end := 0, len(arr)-1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == k {
			return true
		}
		if arr[start] == arr[end] && arr[mid] == arr[start] {
			start++
			end--
		} else {
			if arr[mid] > k {
				if arr[start] <= arr[mid] && isInRange(arr[start], k, arr[mid]) {
					end = mid - 1
				} else {
					start = mid + 1
				}
			} else {
				if arr[mid] <= arr[end] && isInRange(arr[mid], k, arr[end]) {
					start = mid + 1
				} else {
					end = mid - 1
				}
			}
		}
	}
	return false
}

func FindMinimumInRotatedArray(arr []int) int {
	start, end := 0, len(arr)-1
	min := math.MaxInt
	for start <= end {
		mid := start + (end-start)/2
		if min > arr[mid] {
			min = arr[mid]
		}
		if arr[start] <= arr[end] {
			if min > arr[start] {
				return arr[start]
			}
			return min
		}
		if arr[start] <= arr[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return min
}
