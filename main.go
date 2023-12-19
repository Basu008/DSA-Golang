package main

import (
	"fmt"
	"main/binarysearch"
)

func main() {
	fmt.Println("Max weight for shipment", binarysearch.ShipmentOfPackages([]int{5, 4, 5, 2, 3, 4, 5, 6}, 5))
}
