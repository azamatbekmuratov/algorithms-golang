package main

import "fmt"

func quickSort(arr []int, low, high int) {
	if low < high {
		var pivot = partition(arr, low, high)
		quickSort(arr, low, pivot)
		quickSort(arr, pivot+1, high)
	}
}

func partition(arr []int, low, high int) int {
	var pivot = arr[low]
	var i = low
	var j = high

	for i < j {
		for arr[i] <= pivot && i < high {
			i++
		}
		for arr[j] > pivot && j > low {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[low] = arr[j]
	arr[j] = pivot
	return j
}

func main() {
	var arr = []int{84, 72, 72, 72, 84, 72}
	fmt.Printf("Before sorting: %v", arr)

	quickSort(arr, 0, len(arr)-1)
	fmt.Printf("\nAfter sorting: %v", arr)
}
