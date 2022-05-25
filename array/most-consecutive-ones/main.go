//Given a binary array nums, return the maximum number of consecutive 1's in the array.
package main

import (
	"fmt"
)

func findMaxConsecutiveOnes(input []int) int {
	var maxCount int = 0
	var count int = 0
	for _, v := range input {
		if v == 1 {
			count = count + 1
		}
		if v == 0 {
			if count > maxCount {
				maxCount = count
			}
			count = 0
		}
	}
	if count > maxCount {
		maxCount = count
	}
	return maxCount
}

func main() {
	var data []int = []int{1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(data))

}
