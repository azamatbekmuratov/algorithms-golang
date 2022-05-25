package main

import "fmt"

func removeDuplicates(nums *[]int) int {
	if len((*nums)) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len((*nums)); j++ {
		if (*nums)[i] != (*nums)[j] {
			(*nums)[i+1] = (*nums)[j]
			i++
		}
	}

	// fmt.Println(nums)

	return i + 1
}

func main() {
	input := []int{1, 1, 2}
	fmt.Println(removeDuplicates(&input))
	fmt.Println(input)
}
