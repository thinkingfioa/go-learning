package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50}

	newSlice := slice[2:5:5]

	newSlice = append(newSlice, 60, 70)
	fmt.Println("Slice :")
	for index, value := range slice {
		fmt.Println("Index %d, Value: %d", index, value)
	}
	fmt.Println("NewSlice :")
	for index, value := range newSlice {
		fmt.Println("Index %d, Value: %d", index, value)
	}
}
