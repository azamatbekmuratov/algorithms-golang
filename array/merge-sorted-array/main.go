package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := 2
	p2 := len(nums1) - 1
	p3 := len(nums2) - 1

	for i := len(nums1) - 1; i > 2; i-- {
		if nums2[p3] > nums1[p1] {
			nums1[p2] = nums2[p3]
			p2--
			if p3 > 0 {
				p3--
			}
		}
		if nums2[p3] <= nums1[p1] {
			nums1[p2] = nums1[p1]
			p2--
			p1--
		}
	}
	fmt.Println(nums1)
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
}
