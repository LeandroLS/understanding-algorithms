package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[0]
	lower := []int{}
	for i := range arr {
		if arr[i] < pivot {
			lower = append(lower, arr[i])
		}
	}
	higher := []int{}
	for i := range arr {
		if arr[i] > pivot {
			higher = append(higher, arr[i])
		}
	}
	lowerSorted := quickSort(lower)
	higherSorted := quickSort(higher)
	lowerAndPivot := append(lowerSorted, []int{pivot}[0])
	lowerAndPivot = append(lowerAndPivot, higherSorted...)
	return lowerAndPivot
}

func main() {
	fmt.Println(quickSort([]int{50, 78, 1, 4, 8, 90, 102, 3}))
}
