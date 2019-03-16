package main

import "fmt"

func main() {
	var array1 [3]*string

	array2 := [3]*string{new(string), new(string), new(string)}
	*array2[0] = "Red"
	*array2[1] = "Blue"
	*array2[2] = "Green"

	array1 = array2

	*array1[0] = "Yellow"
	// 遍历
	for index := range array1 {
		fmt.Println(*array1[index])
	}

	// 遍历
	for index := range array2 {
		fmt.Println(*array1[index])
	}
}
