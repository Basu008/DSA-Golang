package recursion

import "fmt"

var count int = 0

func CountTo3() {
	fmt.Println(count)
	if count == 3 {
		return
	}
	count++
	CountTo3()
}

func PrintNameNtimes(i int, n int) {
	if i > n {
		return
	}
	fmt.Println("Basu")
	PrintNameNtimes(i+1, n)
}

func Print1toN(i int, n int) {
	if i > n {
		return
	}
	fmt.Println(i)
	Print1toN(i+1, n)
}

func PrintNto1(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	PrintNto1(n - 1)
}

func Print1toN_BackTracking(i int, n int) {
	if i < 1 {
		return
	}
	Print1toN_BackTracking(i-1, n)
	fmt.Println(i)
}

func PrintNto1_BackTracking(i int, n int) {
	if i > n {
		return
	}
	PrintNto1_BackTracking(i+1, n)
	fmt.Println(i)
}

func SumFrom1toN(n int) int {
	if n == 0 {
		return 0
	}
	return n + SumFrom1toN(n-1)
}

func Factorial(n int) int {
	if n < 0 {
		return -1
	}
	if n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func Reverse(input []int, start int, end int) {
	if start >= end {
		return
	}
	input[start], input[end] = input[end], input[start]
	Reverse(input, start+1, end-1)
}

// With only one variable
func ReverseWithOneVar(arr []int, i int) []int {
	if i >= len(arr)/2 {
		return arr
	}
	arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	return ReverseWithOneVar(arr, i+1)
}

func IsPalindrome(input string, i int) bool {
	if i >= len(input)/2 {
		return true
	}
	if input[i] != input[len(input)-i-1] {
		return false
	}
	return IsPalindrome(input, i+1)
}

func Fibonnaci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonnaci(n-1) + Fibonnaci(n-2)
}
