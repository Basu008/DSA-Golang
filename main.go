package main

import (
	"fmt"
	"main/binarysearch"
)

func main() {
	arr := []int{1, 1, 3, 5}
	fmt.Println("Answer", binarysearch.SearchSingleElement(arr))
}
