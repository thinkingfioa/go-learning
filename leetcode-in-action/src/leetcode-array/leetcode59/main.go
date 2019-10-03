package main

import (
	"fmt"
)

func generateMatrix(n int) [][]int {
	num := 1
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	for i := 0; i < (n+1)/2; i++ {
		leftToRight(result, i, i, n-1-i, &num)
		rightToDown(result, n-1-i, i+1, n-1-i, &num)
		downToLeft(result, n-1-i, n-1-i-1, i, &num)
		toLeftUp(result, i, n-1-i-1, i+1, &num)
	}
	return result
}

func leftToRight(result [][]int, row int, start int, end int, num *int) {
	for i := start; i <= end; i++ {
		result[row][i] = *num
		*num = *num + 1
	}
}

func rightToDown(result [][]int, column int, start int, end int, num *int) {
	for i := start; i <= end; i++ {
		result[i][column] = *num
		*num = *num + 1
	}
}

func downToLeft(result [][]int, row int, start int, end int, num *int) {
	for i := start; i >= end; i-- {
		result[row][i] = *num
		*num = *num + 1
	}
}

func toLeftUp(result [][]int, column int, start int, end int, num *int) {
	for i := start; i >= end; i-- {
		result[i][column] = *num
		*num = *num + 1
	}
}

func main() {
	n := 3
	result := generateMatrix(n)
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			fmt.Printf("%d, ", result[i][j])
		}
		fmt.Println()
	}
}
