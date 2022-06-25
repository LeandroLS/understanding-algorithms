package main

import "fmt"

func main() {
	telephoneList := make(map[string]string)
	telephoneList["Esther"] = "111111111"
	telephoneList["Ben"] = "222222222"
	for i := range telephoneList {
		fmt.Printf("Name: %s. Number: %s \n", i, telephoneList[i])
	}
}
