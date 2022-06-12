package main

import (
	"fmt"
)

func searchLower(arr []int) int {
	lower := arr[0]
	lowerIndex := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < lower {
			lower = arr[i]
			lowerIndex = i
		}
	}
	return lowerIndex
}

func selectionSort(arr []int) []int {
	var newArr []int
	for len(arr) >= 1 {
		lower := searchLower(arr)
		newArr = append(newArr, arr[lower])
		arr = append(arr[:lower], arr[lower+1:]...)
	}
	return newArr
}

func main() {
	fmt.Println(selectionSort([]int{5, 4, 3, 2, 1}))
}
