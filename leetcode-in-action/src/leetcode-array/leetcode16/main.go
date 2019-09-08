package main

import (
	"fmt"
	"sort"
)

const INT_MAX = int(^uint(0) >> 1)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := INT_MAX
	for i := 0; i < len(nums); i++ {
		p1 := i + 1
		p2 := len(nums) - 1
		for p1 < p2 {
			calNum := nums[i] + nums[p1] + nums[p2]
			if result == INT_MAX || intAbs(calNum-target) < intAbs(result-target) {
				result = calNum
			}
			if result == INT_MAX || calNum-target < 0 {
				p1++
			} else if calNum-target > 0 {
				p2--
			} else {
				break
			}
		}
	}
	return result
}

func intAbs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	result := threeSumClosest(nums, target)
	fmt.Println(result)
}
