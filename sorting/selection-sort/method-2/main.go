package main

import "fmt"

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIdx := 0
	for idx, el := range arr {
		if el < smallest {
			smallest = el
			smallestIdx = idx
		}
	}
	return smallestIdx
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func selectionSort(arr []int) []int {
	var newArr = []int{}
	high := len(arr)
	for i := 0; i < high; i++ {
		smallest := findSmallest(arr)
		newArr = append(newArr, arr[smallest])
		arr = remove(arr, smallest)
	}
	return newArr
}

func main() {
	var input = []int{5, 3, 6, 2, 10}
	fmt.Println(selectionSort(input))
}
