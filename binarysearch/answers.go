package binarysearch

func GetSqrtOfNumber(n int) int {
	start, end := 1, n
	ans := -1
	for start <= end {
		mid := start + (end-start)/2
		if mid*mid <= n {
			ans = mid
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return ans
}

func NthRootOfANumber(m int, n int) int {
	start, end := 1, m
	currentAns := 1
	for start <= end {
		mid := start + (end-start)/2
		ans := Pow(m, n, mid)
		if ans == 1 {
			return mid
		}
		if ans == 2 {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return currentAns
}

func Pow(x int, y int, mid int) int {
	ans := 1
	for i := 1; i <= y; i++ {
		ans *= mid
		if ans > x {
			return 2
		}
	}
	if ans == x {
		return 1
	}
	return 0
}
