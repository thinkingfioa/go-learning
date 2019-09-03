package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// key is num, value is index
	markMap := make(map[int]int)
	for index := 0; index < len(nums); index++ {
		suples := target - nums[index]
		value, ok := markMap[suples]
		if ok && value != index {
			return []int{value, index}
		}
		markMap[nums[index]] = index
	}
	return nil
}

func main() {
	nums := []int{3, 3}
	target := 6
	result := twoSum(nums, target)
	for index := 0; index < len(result); index++ {
		fmt.Printf("%d, ", result[index])
	}
}
