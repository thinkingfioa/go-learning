package main

import "fmt"

func searchRange(nums []int, target int) []int {
	result := make([]int, 0)
	leftIndex := leftBinarySearch(nums, 0, len(nums)-1, target)
	if leftIndex == -1 {
		return append(result, -1, -1)
	}
	result = append(result, leftIndex)
	rightIndex := rightBinarySearch(nums, leftIndex, len(nums)-1, target)
	result = append(result, rightIndex)
	return result
}

func leftBinarySearch(nums []int, left int, right int, target int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	if (mid == left || nums[mid-1] != target) && (nums[mid] == target) {
		return mid
	}
	if nums[mid] >= target {
		return leftBinarySearch(nums, left, mid-1, target)
	} else {
		return leftBinarySearch(nums, mid+1, right, target)
	}
}

func rightBinarySearch(nums []int, left int, right int, target int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	if (mid == right || nums[mid+1] != target) && (nums[mid] == target) {
		return mid
	}
	if nums[mid] > target {
		return rightBinarySearch(nums, left, mid-1, target)
	} else {
		return rightBinarySearch(nums, mid+1, right, target)
	}
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	result := searchRange(nums, target)
	for i := 0; i < len(result); i++ {
		fmt.Printf("%d, ", result[i])
	}
}
