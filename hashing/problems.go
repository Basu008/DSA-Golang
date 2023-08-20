package hashing

import (
	"fmt"
)

func OccuranceOfNumbersInArray(arr []int) {
	hashMap := make(map[int]int)
	for _, data := range arr {
		hashMap[data]++
	}
	for key, value := range hashMap {
		fmt.Println(key, "occured", value, "times")
	}
}

func HighestAndLowestFrequenct(arr []int) {
	hashMap := make(map[int]int)
	for _, data := range arr {
		hashMap[data]++
	}
	maxCount := 0
	minCount := 123123123123
	var highest int = arr[0]
	var lowest int = arr[0]
	fmt.Println(hashMap)
	for key, value := range hashMap {
		if maxCount < value {
			maxCount = value
			highest = key
		}
		if minCount > value {
			minCount = value
			lowest = key
		}
	}
	fmt.Println("Highest value : ", highest)
	fmt.Println("Lowest value : ", lowest)
}
