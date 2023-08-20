package sorting

import "fmt"

func SelectionSort(arr []int) {
	size := len(arr)
	fmt.Println("Before sorting", arr)
	for i := 0; i < size; i++ {
		min := arr[i]
		pos := i
		for j := i; j < size; j++ {
			if min > arr[j] {
				min = arr[j]
				pos = j
			}
		}
		arr[i], arr[pos] = arr[pos], arr[i]
	}
	fmt.Println("After sorting", arr)
}

func BubbleSort(arr []int) {
	size := len(arr)
	fmt.Println("Before sorting", arr)
	for i := 0; i < size-1; i++ {
		for j := 0; j < size-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("After sorting", arr)
}

func RecursiveBubbleSort(arr []int, n int) {
	if n == 0 {
		return
	}
	for j := 0; j < n-1; j++ {
		if arr[j] > arr[j+1] {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
	RecursiveBubbleSort(arr, n-1)
}

func InsertionSort(arr []int) {
	size := len(arr)
	fmt.Println("Before sorting", arr)
	for i := 1; i < size-1; i++ {
		for j := i; j >= 0; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	fmt.Println("After sorting", arr)
}

func RecursiveInsertionSort(arr []int, i int, n int) {
	if i == n {
		return
	}
	for j := i; j > 0; j-- {
		if arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	RecursiveInsertionSort(arr, i+1, len(arr))
}

func QuickSort(arr []int, start int, end int) {
	if start > end {
		return
	}
	pivot := getPivot(arr, start, end)
	QuickSort(arr, start, pivot-1)
	QuickSort(arr, pivot+1, end)
}

func getPivot(arr []int, start int, end int) int {
	if start > end {
		return end
	}
	pivot := arr[end]
	pivotIndex := start
	for i := start; i < end; i++ {
		if arr[i] < pivot {
			arr[i], arr[pivotIndex] = arr[pivotIndex], arr[i]
			pivotIndex++
		}
	}
	arr[end], arr[pivotIndex] = arr[pivotIndex], arr[end]
	return pivotIndex
}

func MergeSort(arr []int, start int, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	MergeSort(arr, start, mid)
	MergeSort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

func merge(arr []int, start int, mid int, end int) {
	i := start
	j := mid + 1
	fmt.Println("Left:", arr[start:mid+1])
	fmt.Println("Right:", arr[mid+1:end+1])
	newArr := []int{}
	for i <= mid && j <= end {
		if arr[i] <= arr[j] {
			newArr = append(newArr, arr[i])
			i++
		} else {
			newArr = append(newArr, arr[j])
			j++
		}
	}
	for i <= mid {
		newArr = append(newArr, arr[i])
		i++
	}
	for j <= end {
		newArr = append(newArr, arr[j])
		j++
	}
	fmt.Println("Sorted :", newArr)
	for i := start; i <= end; i++ {
		arr[i] = newArr[i-start]
	}
}
