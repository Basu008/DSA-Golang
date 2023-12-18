package binarysearch

import "math"

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

func BananasPerHour(arr []int, hours int) int {
	max := math.MinInt
	for _, stack := range arr {
		if stack >= max {
			max = stack
		}
	}
	ans := 0
	start, end := 1, max
	for start <= end {
		mid := start + (end-start)/2
		totalHours := CalculateBanana(mid, arr)
		if totalHours <= hours {
			if totalHours == hours {
				ans = mid
			}
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return ans
}

func CalculateBanana(banana int, stacks []int) int {
	totalHours := 0
	for _, stack := range stacks {
		hourForThisStack := int(math.Ceil(float64(stack) / float64(banana)))
		totalHours += hourForThisStack
	}
	return totalHours
}

func CalculateBouquet(flowers []int, bCount int, fCount int) int {
	max, min := math.MinInt, math.MaxInt
	for _, flowerBloomDay := range flowers {
		if flowerBloomDay <= min {
			min = flowerBloomDay
		}
		if flowerBloomDay >= max {
			max = flowerBloomDay
		}
	}
	ans := -1
	start, end := min, max
	for start <= end {
		mid := start + (end-start)/2
		if canBouquetBeFormed(flowers, mid, bCount, fCount) {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return ans
}

func canBouquetBeFormed(arr []int, day int, bCount int, fCount int) bool {
	count := 0
	bouquetFormed := 0
	for _, dayToBloom := range arr {
		if dayToBloom <= day {
			count += 1
		} else {
			b := count / fCount
			bouquetFormed += b
			count = 0
		}
	}
	b := count / fCount
	bouquetFormed += b
	count = 0
	return bouquetFormed >= bCount
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