package main

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	fmt.Println(nums)
	return i
}

func main() {
	input := []int{3, 2, 2, 3}
	fmt.Println(removeElement(input, 3))
}
