package main

import "fmt"

func minPathSum(grid [][]int) int {
	row := len(grid)
	column := len(grid[0])
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, column)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < column; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < row; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < row; i++ {
		for j := 1; j < column; j++ {
			dp[i][j] = minVal(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[row-1][column-1]
}

func minVal(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Printf("%d\n", minPathSum(grid))
}
