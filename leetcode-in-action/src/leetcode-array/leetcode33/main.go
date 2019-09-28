package main

import "fmt"

func search(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, left int, right int, target int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	if nums[mid] == target {
		return mid
	}

	if nums[mid] >= nums[left] {
		// left sorted
		if target >= nums[left] && target < nums[mid] {
			return binarySearch(nums, left, mid-1, target)
		} else {
			return binarySearch(nums, mid+1, right, target)
		}
	} else {
		// right sorted
		if target <= nums[right] && target > nums[mid] {
			return binarySearch(nums, mid+1, right, target)
		} else {
			return binarySearch(nums, left, mid-1, target)
		}
	}
}

func main() {
	//nums := []int{4, 5, 6, 7, 0, 1, 2}
	//target := 0

	nums := []int{5, 1, 3}
	target := 5

	fmt.Printf("%d", search(nums, target))
}
