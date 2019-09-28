package main

import "fmt"

func searchInsert(nums []int, target int) int {
	return binarySearchForWhile(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, left int, right int, target int) int {
	if left > right {
		return left
	}
	mid := (left + right) / 2
	if nums[mid] == target {
		return mid
	}
	if nums[mid] > target {
		return binarySearch(nums, left, mid-1, target)
	} else {
		return binarySearch(nums, mid+1, right, target)
	}
}

func binarySearchForWhile(nums []int, left int, right int, target int) int {
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Printf("%d", searchInsert(nums, target))
}
