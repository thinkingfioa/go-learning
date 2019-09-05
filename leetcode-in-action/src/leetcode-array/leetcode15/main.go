package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		p1 := i + 1
		p2 := len(nums) - 1
		target := -nums[i]
		for p1 < len(nums) && p1 < p2 {
			if nums[p1]+nums[p2] == target {
				oneResult := []int{nums[i], nums[p1], nums[p2]}
				result = append(result, oneResult)
				p1++
				p2--
				for p1 < p2 && nums[p1] == nums[p1-1] {
					p1++
				}
				for p1 < p2 && nums[p2] == nums[p2+1] {
					p2--
				}
			} else if nums[p1]+nums[p2] > target {
				p2--
			} else {
				p1++
			}
		}
	}
	return result
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			fmt.Printf("%d, ", result[i][j])
		}
		fmt.Println()
	}
}
