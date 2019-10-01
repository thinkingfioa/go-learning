package main

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix)
	// 水平交换，交换公式 (i, j) ---- (i, n-i-j)
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			swap(matrix, i, j, i, n-1-j)
		}
	}
	// 对角线交换，交换公式 (i,j) ---- (n-1-j, n-1-i)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			swap(matrix, i, j, n-1-j, n-1-i)
		}
	}
}

func swap(matrix [][]int, x1 int, y1 int, x2 int, y2 int) {
	temp := matrix[x1][y1]
	matrix[x1][y1] = matrix[x2][y2]
	matrix[x2][y2] = temp
}

func main() {
	matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(matrix)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d, ", matrix[i][j])
		}
		fmt.Println()
	}
}

func print(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d, ", matrix[i][j])
		}
		fmt.Println()
	}
}
