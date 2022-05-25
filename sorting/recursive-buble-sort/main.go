package main

import "fmt"

func bubleSort(arr []int, size int) []int {
	if size == 1 {
		return arr
	}

	var i = 0
	for i < size-1 {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
		i++
	}

	bubleSort(arr, size-1)

	return arr
}

func main() {
	var n = []int{84, 72, 72, 72, 84, 72}
	fmt.Println(bubleSort(n, len(n)))
}
