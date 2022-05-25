package main

import "fmt"

type result struct {
	found bool
	index int
}

func binarySearch(list []int, item int) result {
	var low int = 0
	var high int = len(list) - 1

	for low <= high {
		mid := (low + high)
		guess := list[mid]
		if guess == item {
			return result{found: true, index: mid}
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return result{found: false}
}

func main() {
	var list = []int{1, 3, 5, 7, 9}
	fmt.Println(binarySearch(list, 5))
}
