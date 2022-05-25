package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	less := []int{}
	greater := []int{}

	for i := 1; i < len(arr); i++ {
		if arr[i] <= pivot {
			less = append(less, arr[i])
		} else if arr[i] > pivot {
			greater = append(greater, arr[i])
		}
	}
	leftSide := quickSort(less)
	rightSide := append([]int{pivot}, quickSort(greater)...)
	result := append(leftSide, rightSide...)

	return result
}

func main() {
	input := []int{10, 5, 2, 3}
	fmt.Println(quickSort(input))
}
