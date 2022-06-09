package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	orderedList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result, err := binarySearch(orderedList, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

}

func binarySearch(orderedList []int, itemToSearch int) (int, error) {
	low := 0
	high := len(orderedList) - 1
	for low <= high {
		middle := (low + high) / 2
		guess := orderedList[middle]
		if guess == itemToSearch {
			return middle, nil
		} else if guess > itemToSearch {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}

	return 0, errors.New("Item dont found")
}
