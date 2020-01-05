package main

import "fmt"

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}

func main() {
	//nums := []int{4, 5, 6, 7, 0, 1, 2}
	nums := []int{1, 2}
	fmt.Print(findMin(nums))
}
