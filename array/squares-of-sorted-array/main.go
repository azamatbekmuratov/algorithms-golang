package main

import (
	"fmt"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	var left int = 0
	var right int = n - 1

	for i := n - 1; i >= 0; i-- {
		var square int
		if Abs(nums[left]) < Abs(nums[right]) {
			square = nums[right]
			right--
		} else {
			square = nums[left]
			left++
		}
		result[i] = square * square
	}

	return result
}

func main() {
	input := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(input))
}
