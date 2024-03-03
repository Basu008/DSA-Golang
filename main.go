package main

import (
	"fmt"

	"github.com/Basu008/DSA-Golang/graph"
)

func main() {
	// input := graph.GraphInput{
	// 	Nodes:    5,
	// 	Edges:    [][]int{{0, 1, 5}, {0, 3, 2}, {3, 1, 2}, {1, 2, 5}, {1, 4, 1}, {4, 2, 1}},
	// 	Directed: true,
	// }
	// fmt.Printf("MST? %v", graph.CreateGraph(input).Bridges())
	// matrix := [][]int{
	// 	{1, 1, 1, 1},
	// 	{1, 1, 0, 1},
	// 	{1, 1, 1, 1},
	// 	{1, 1, 0, 0},
	// 	{1, 0, 0, 0},
	// }
	arr := []int{2, 5, 7}
	fmt.Println("Min multiplication :", graph.MinimumMultiplication(3, 30, arr))
}
