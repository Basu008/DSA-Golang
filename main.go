package main

import (
	"fmt"
	"main/binarysearch"
)

func main() {
	fmt.Println("Guldasta", binarysearch.CalculateBouquet([]int{7, 7, 7, 7, 13, 11, 12, 7}, 2, 3))
}
