package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	rowNum := len(matrix)
	if rowNum == 0 {
		return false
	}
	columnNum := len(matrix[0])
	if columnNum == 0 {
		return false
	}
	if matrix[0][0] > target {
		return false
	}
	if matrix[rowNum-1][columnNum-1] < target {
		return false
	}
	rowIndex := binaryColumnSearch(matrix, 0, 0, rowNum-1, target)
	fmt.Printf("rowIndex:%d\n", rowIndex)
	findIndex := binaryRowSearch(matrix, rowIndex, 0, columnNum-1, target)
	fmt.Printf("findIndex:%d\n", findIndex)
	return matrix[rowIndex][findIndex] == target
}

func binaryRowSearch(matrix [][]int, rowIndex int, left int, right int, target int) int {
	for left <= right {
		mid := (left + right) / 2
		if target == matrix[rowIndex][mid] {
			return mid
		} else if target < matrix[rowIndex][mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left - 1
}

func binaryColumnSearch(matrix [][]int, columnIndex int, left int, right int, target int) int {
	for left <= right {
		mid := (left + right) / 2
		if target == matrix[mid][columnIndex] {
			return mid
		} else if target < matrix[mid][columnIndex] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left - 1
}

func main() {
	//matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	//target := 3

	matrix := [][]int{{1}, {3}, {5}}
	target := 3

	fmt.Println(searchMatrix(matrix, target))
}
