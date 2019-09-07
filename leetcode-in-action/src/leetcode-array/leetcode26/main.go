package main

import "fmt"

func removeDuplicates(nums []int) int {
	p1 := -1
	for p2 := 0; p2 < len(nums); p2++ {
		if p1 == -1 {
			p1++
			nums[p1] = nums[p2]
		} else {
			if nums[p1] != nums[p2] {
				p1++
				nums[p1] = nums[p2]
			}
		}
	}
	return p1 + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	result := removeDuplicates(nums)
	for index := 0; index < result; index++ {
		fmt.Printf("%d, ", nums[index])
	}
}
