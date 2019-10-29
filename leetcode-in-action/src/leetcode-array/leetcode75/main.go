package main

import "fmt"

func sortColors(nums []int) {
	left := 0
	right := len(nums) - 1
	for left <= right && nums[left] == 0 {
		left++
	}
	for left <= right && nums[right] == 2 {
		right--
	}
	move := left
	for move <= right {
		if nums[move] == 0 {
			swap(nums, left, move)
			move++
			for left < move && nums[left] == 0 {
				left++
			}
		} else if nums[move] == 2 {
			swap(nums, move, right)
			for right > move && nums[right] == 2 {
				right--
			}
		} else {
			move++
		}
	}
}

func swap(nums []int, aIndex int, bIndex int) {
	temp := nums[aIndex]
	nums[aIndex] = nums[bIndex]
	nums[bIndex] = temp
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0, 0}
	sortColors(nums)

	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d, ", nums[i])
	}
}
