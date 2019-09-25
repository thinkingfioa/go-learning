package main

import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int) {
	move := len(nums) - 1
	for ; move >= 1; move-- {
		if nums[move] > nums[move-1] {
			break
		}
	}
	if move == 0 {
		sort.Ints(nums)
		return
	}
	move = move - 1
	posIndex := len(nums) - 1
	for ; posIndex > 0; posIndex-- {
		if nums[posIndex] > nums[move] {
			break
		}
	}
	// 交换
	swap(nums, move, posIndex)
	p1 := move + 1
	p2 := len(nums) - 1
	for p1 < p2 {
		swap(nums, p1, p2)
		p1++
		p2--
	}
}

func swap(nums []int, one int, two int) {
	temp := nums[one]
	nums[one] = nums[two]
	nums[two] = temp
}

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	for index := 0; index < len(nums); index++ {
		fmt.Printf("%d, ", nums[index])
	}
}
