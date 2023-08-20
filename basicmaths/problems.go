package basicmaths

import (
	"fmt"
	"math"
)

func CountDigits(input int) int {
	var length int
	number := input
	for number > 0 {
		number = number / 10
		length++
	}
	return length
}

func ReverseNumber(input int) int {
	var reversedNumber int
	number := input
	for number > 0 {
		reversedNumber = reversedNumber*10 + (number % 10)
		number /= 10
	}
	return reversedNumber
}

func IsNumberPalindrom(input int) bool {
	var reversedNumber = ReverseNumber(input)
	return reversedNumber == input
}

func CalculateGCD(a int, b int) int {
	if b == 0 {
		return a
	}
	return CalculateGCD(b, a%b)
}

func IsNumberArmstrong(input int) bool {
	digits := CountDigits(input)
	var sum float64
	num := input
	for num != 0 {
		sum += math.Pow(float64(num%10), float64(digits))
		num /= 10
	}
	return sum == float64(input)
}

func AllPossibleDivisor(n int) {
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			fmt.Print(i)
			fmt.Print(",")
			if i != n/i {
				fmt.Print(n / i)
				fmt.Print(",")
			}
		}
	}
}

func IsNumberPrime(n int) bool {
	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
