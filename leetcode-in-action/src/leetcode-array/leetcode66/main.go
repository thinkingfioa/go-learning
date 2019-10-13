package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	result := make([]int, 0)
	surplus := 1
	for i := len(digits) - 1; i >= 0; i-- {
		d := (surplus + digits[i]) % 10
		surplus = (surplus + digits[i]) / 10
		result = append(result, d)
	}
	for surplus != 0 {
		d := surplus % 10
		surplus = surplus / 10
		result = append(result, d)
	}
	return reverse(result)
}

func reverse(result []int) []int {
	length := len(result)
	for i := 0; i < length/2; i++ {
		swap(result, i, length-1-i)
	}
	return result
}

func swap(slice []int, a int, b int) {
	temp := slice[a]
	slice[a] = slice[b]
	slice[b] = temp
}

func main() {
	digits := []int{4, 3, 2, 2}
	result := plusOne(digits)

	for i := 0; i < len(result); i++ {
		fmt.Printf("%d", result[i])
	}
}
