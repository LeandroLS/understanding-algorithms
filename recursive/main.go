package main

import (
	"fmt"
)

func soma(arr []int) int {
	arrLenght := len(arr)
	if arrLenght == 1 {
		return arr[0]
	}
	arrReduced := arr[:arrLenght-1]
	value := arr[arrLenght-1]
	return value + soma(arrReduced)
}

func countItens(arr []int) int {
	arrReduced := arr[:len(arr)-1]
	if len(arrReduced) >= 1 {
		return 1 + countItens(arrReduced)
	} else {
		return 1
	}
}

func higherValue(arr []int) int {
	arrReduced := arr[:len(arr)-1]
	value := arr[len(arr)-1]
	if len(arrReduced) >= 1 {
		expectedToBeHigher := value
		value := higherValue(arrReduced)
		if value >= expectedToBeHigher {
			return value
		} else {
			return expectedToBeHigher
		}
	} else {
		return 0
	}
}

func main() {
	toSumAndCount := []int{1, 2, 15, 3, 4, 5, 6}
	fmt.Println("Sum:", soma(toSumAndCount))
	fmt.Println("Count itens:", countItens(toSumAndCount))
	fmt.Println("Higher value:", higherValue(toSumAndCount))
}
