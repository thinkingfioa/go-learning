package main

import "fmt"

func maxArea(height []int) int {
	maxArea := 0
	left := 0
	right := len(height) - 1
	for left < right {
		area := (right - left) * min(height[left], height[right])
		maxArea = max(maxArea, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func min(num1 int, num2 int) int {
	if num1 < num2 {
		return num1
	}

	return num2
}

func max(num1 int, num2 int) int {
	if num1 < num2 {
		return num2
	}
	return num1
}

func main() {
	//height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	height := []int{1, 2, 4, 3}
	fmt.Println(maxArea(height))
}
