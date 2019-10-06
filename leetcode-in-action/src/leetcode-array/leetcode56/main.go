package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	for i := 0; i < len(intervals); i++ {
		sort.Ints(intervals[i])
	}
	quickSort(intervals, 0, len(intervals)-1)
	result := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		if i < len(intervals)-1 && intervals[i][1] >= intervals[i+1][0] {
			maxV := intervals[i][1]
			end := i + 1
			for ; end < len(intervals); end++ {
				if intervals[end][0] <= maxV {
					maxV = maxVal(intervals[end][1], maxV)
				} else {
					break
				}
			}
			oneResult := make([]int, 0)
			oneResult = append(oneResult, intervals[i][0])
			oneResult = append(oneResult, maxV)
			result = append(result, oneResult)
			i = end - 1
		} else {
			result = append(result, intervals[i])
		}
	}
	return result
}

func quickSort(intervals [][]int, left int, right int) {
	if left >= right {
		return
	}
	index := quickSort0(intervals, left, right)
	quickSort(intervals, left, index-1)
	quickSort(intervals, index+1, right)
}

func quickSort0(intervals [][]int, left int, right int) int {
	if left >= right {
		return -1
	}
	swapVal := intervals[left]
	for left < right {

		for left < right && swapVal[0] <= intervals[right][0] {
			right = right - 1
		}
		intervals[left] = intervals[right]
		for left < right && swapVal[0] > intervals[left][0] {
			left = left + 1
		}
		intervals[right] = intervals[left]
	}
	intervals[left] = swapVal
	return left
}

func maxVal(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	//intervals := [][]int{{1, 4}, {4, 5}}
	//intervals := [][]int{{1, 10}, {2, 4}, {4, 5}}
	result := merge(intervals)
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			fmt.Printf("%d, ", result[i][j])
		}
		fmt.Println()
	}
}
