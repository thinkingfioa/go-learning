package main

import "fmt"

func removeElement(nums []int, val int) int {
	p1 := 0
	for ; p1 < len(nums); p1++ {
		if nums[p1] == val {
			p2 := p1 + 1
			for ; p2 < len(nums); p2++ {
				if nums[p2] != val {
					swap(nums, p1, p2)
					break
				}
			}
			if p2 >= len(nums) {
				break
			}
		}
	}
	return p1
}

func swap(nums []int, iOne int, iTwo int) {
	temp := nums[iOne]
	nums[iOne] = nums[iTwo]
	nums[iTwo] = temp
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	len := removeElement(nums, val)
	for i := 0; i < len; i++ {
		fmt.Printf("%d, ", nums[i])
	}
}
