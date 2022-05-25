package main

import "fmt"

func moveZeros(nums []int) {
	p1 := 0
	p2 := 1
	for p2 < len(nums) {
		if nums[p1] == 0 && nums[p2] != 0 {
			nums[p1] = nums[p2]
			nums[p2] = 0
			p1++
		} else if nums[p1] == 0 && nums[p2] == 0 {
			p2++
		} else {
			p1++
			p2++
		}
	}
	fmt.Println(nums)
}

func main() {
	var input []int = []int{0, 1, 0, 3, 12}
	moveZeros(input)
}
