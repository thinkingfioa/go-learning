package main

import (
	"fmt"
)

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func maxSubArray(nums []int) int {
	return subArray(nums, 0, len(nums)-1)
}

func subArray(nums []int, beginIndex int, endIndex int) int {
	if beginIndex > endIndex {
		return INT_MIN
	}
	if beginIndex == endIndex {
		return nums[beginIndex]
	}
	mid := (beginIndex + endIndex) / 2
	leftMax := subArray(nums, beginIndex, mid-1)
	rightMax := subArray(nums, mid+1, endIndex)
	// [left, mid] maxValue
	maxValue := INT_MIN
	sum := 0
	for i := mid; i >= beginIndex; i-- {
		sum += nums[i]
		if maxValue < sum {
			maxValue = sum
		}
	}
	sum = maxValue
	for i := mid + 1; i <= endIndex; i++ {
		sum += nums[i]
		if maxValue < sum {
			maxValue = sum
		}
	}
	return max(maxValue, max(leftMax, rightMax))
}

func max(v1 int, v2 int) int {
	if v1 > v2 {
		return v1
	} else {
		return v2
	}
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
