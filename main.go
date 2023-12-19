package main

import (
	"fmt"
	"main/binarysearch"
)

func main() {
	fmt.Println("Minimum", binarysearch.FindSmallestDivisor([]int{8, 4, 2, 3}, 10))
}
