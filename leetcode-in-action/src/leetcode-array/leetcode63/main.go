package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	row := len(obstacleGrid)
	column := len(obstacleGrid[0])
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, column)
	}
	flag := false
	for i := 0; i < row; i++ {
		if flag || obstacleGrid[i][0] != 0 {
			dp[i][0] = 0
			flag = true

		} else {
			dp[i][0] = 1
		}
	}
	flag = false
	for i := 0; i < column; i++ {
		if flag || obstacleGrid[0][i] != 0 {
			dp[0][i] = 0
			flag = true
		} else {
			dp[0][i] = 1
		}
	}
	for i := 1; i < row; i++ {
		for j := 1; j < column; j++ {
			if obstacleGrid[i][j] != 0 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[row-1][column-1]
}

func main() {
	//obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	obstacleGrid := [][]int{{1, 0}}
	fmt.Printf("%d\n", uniquePathsWithObstacles(obstacleGrid))
}
