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
