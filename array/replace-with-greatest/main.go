package main

import "fmt"

func replaceElements(arr []int) []int {
	max := 0
	if len(arr) == 0 {
		return arr
	}

	if len(arr) == 1 {
		arr[0] = -1
		return arr
	}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > max {
				max = arr[j]
			}
		}
		if i == len(arr)-1 {
			arr[i] = -1
		} else {
			arr[i] = max
		}
		max = 0
	}
	return arr
}

func main() {
	input := []int{17, 18, 5, 4, 6, 1}
	fmt.Println(replaceElements(input))
}
