package main

import (
	"fmt"
	"main/binarysearch"
)

func main() {
	arr := []int{4, 5, 6, 7, 0, 1, 2, 3}
	fmt.Println("Minimum element", binarysearch.FindMinimumInRotatedArray(arr))
}
