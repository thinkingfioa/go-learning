package main

import "fmt"

func uniquePaths(m int, n int) int {
	num := 0
	dfs(1, 1, m, n, &num)
	return num
}

func dfs(row int, column int, m int, n int, num *int) {
	if row == n && column == m {
		*num = *num + 1
		return
	}
	if row > n {
		return
	}
	if column > m {
		return
	}
	dfs(row+1, column, m, n, num)
	dfs(row, column+1, m, n, num)
}

func main() {
	//m := 7
	//n := 3
	m := 51
	n := 9
	fmt.Printf("%d\n", uniquePaths(m, n))
}
