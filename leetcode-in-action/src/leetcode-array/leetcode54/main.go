package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	row := len(matrix)
	if row == 0 {
		return make([]int, 0)
	}
	column := len(matrix[0])
	result := make([]int, 0)
	min := minInt(row, column)
	for i := 0; i < (min+1)/2; i++ {
		leftToRight(matrix, i, i, column-1-i, &result)
		rightToDown(matrix, column-1-i, i+1, row-1-i, &result)
		if i != row-1-i {
			downToLeft(matrix, row-1-i, column-1-i-1, i, &result)
		}
		if column-1-i != i {
			leftToUp(matrix, i, row-1-i-1, i+1, &result)
		}
	}
	return result
}

func minInt(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func leftToRight(matrix [][]int, row int, begin int, end int, result *[]int) {
	for i := begin; i <= end; i++ {
		*result = append(*result, matrix[row][i])
	}
}

func rightToDown(matrix [][]int, column int, begin int, end int, result *[]int) {
	for i := begin; i <= end; i++ {
		*result = append(*result, matrix[i][column])
	}
}

func downToLeft(matrix [][]int, row int, begin int, end int, result *[]int) {
	for i := begin; i >= end; i-- {
		*result = append(*result, matrix[row][i])
	}
}

func leftToUp(matrix [][]int, column int, begin int, end int, result *[]int) {
	for i := begin; i >= end; i-- {
		*result = append(*result, matrix[i][column])
	}
}

func main() {
	//matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	//matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	matrix := [][]int{{1}, {2}, {3}, {4}}
	result := spiralOrder(matrix)
	for i := 0; i < len(result); i++ {
		fmt.Printf("%d, ", result[i])
	}
}
