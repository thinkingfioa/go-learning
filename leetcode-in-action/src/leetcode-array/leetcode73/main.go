package main

import "fmt"

func setZeroes(matrix [][]int) {
	rowNum := len(matrix)
	columnNum := len(matrix[0])
	selectRow := -1
	selectColumn := -1
	// row find have 0
	for i := 0; i < rowNum; i++ {
		if haveRowZero(matrix, i) {
			selectRow = i
			break
		}
	}
	if selectRow == -1 {
		return
	}
	// column find have 0
	for i := 0; i < columnNum; i++ {
		if haveColumnZero(matrix, i) {
			selectColumn = i
			break
		}
	}
	// mark 0
	for i := 0; i < rowNum; i++ {
		if i == selectRow {
			continue
		}
		for j := 0; j < columnNum; j++ {
			if j == selectColumn {
				continue
			}
			if matrix[i][j] == 0 {
				matrix[selectRow][j] = 0
				matrix[i][selectColumn] = 0
			}
		}
	}
	// set 0
	for i := 0; i < rowNum; i++ {
		if matrix[i][selectColumn] == 0 {
			if i == selectRow {
				continue
			}
			for j := 0; j < columnNum; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for i := 0; i < columnNum; i++ {
		if matrix[selectRow][i] == 0 {
			if i == selectColumn {
				continue
			}
			for j := 0; j < rowNum; j++ {
				matrix[j][i] = 0
			}
		}
	}

	for i := 0; i < rowNum; i++ {
		matrix[i][selectColumn] = 0
	}
	for i := 0; i < columnNum; i++ {
		matrix[selectRow][i] = 0
	}
}

func haveRowZero(matrix [][]int, row int) bool {
	column := len(matrix[row])
	for i := 0; i < column; i++ {
		if matrix[row][i] == 0 {
			return true
		}
	}
	return false
}

func haveColumnZero(matrix [][]int, column int) bool {
	row := len(matrix)
	for i := 0; i < row; i++ {
		if matrix[i][column] == 0 {
			return true
		}
	}
	return false
}

func main() {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d, ", matrix[i][j])
		}
		fmt.Println()
	}
}
