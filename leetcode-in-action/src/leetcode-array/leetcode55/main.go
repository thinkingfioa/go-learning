package main

import "fmt"

func canJump(nums []int) bool {
	lenVal := len(nums)
	maxD := nums[0]
	if lenVal == 1 {
		return true
	}
	for i := 1; i <= maxD; i++ {
		maxD = maxVal(maxD, i+nums[i])
		if maxD >= lenVal-1 {
			return true
		}
	}
	return false
}

func maxVal(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//nums := []int{3, 2, 1, 0, 4}
	//nums := []int{2, 2, 1, 1, 4}
	nums := []int{0}
	fmt.Println(canJump(nums))
}
