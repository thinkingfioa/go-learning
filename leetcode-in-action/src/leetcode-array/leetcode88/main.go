package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {

	p1 := m - 1
	p2 := n - 1
	move := m + n - 1
	for ; p2 >= 0 && p1 >= 0; move-- {
		if nums1[p1] >= nums2[p2] {
			nums1[move] = nums1[p1]
			p1--
		} else {
			nums1[move] = nums2[p2]
			p2--
		}
	}
	for ; p2 >= 0; move-- {
		nums1[move] = nums2[p2]
		p2--
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	for i := 0; i < m+n; i++ {
		fmt.Printf("%d", nums1[i])
	}
}
