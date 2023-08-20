package main

import (
	"fmt"
	"main/binarysearch"
)

func main() {
	arr := []int{3, 1, 2, 3, 3, 3, 3}
	fmt.Println(1, "found at :", binarysearch.FindInRotatedArray2(arr, 1))
}
