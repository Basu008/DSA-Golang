package main

import (
	"fmt"
	"main/binarysearch"
)

func main() {
	arr := []int{1, 5, 1, 2}
	fmt.Println("Answer", binarysearch.PeakElement(arr))
}
