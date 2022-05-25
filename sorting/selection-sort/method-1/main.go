package main

import "fmt"

func main() {
	sample := []int{3, 4, 5, 2, 1}
	selectionSort(sample)
}

func selectionSort(arr []int) {
	len := len(arr)
	fmt.Println(len - 1)
	for i := 0; i < len-1; i++ {
		minIndex := i
		for j := i + 1; j < len; j++ {
			if arr[j] < arr[minIndex] {
				arr[j], arr[minIndex] = arr[minIndex], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
