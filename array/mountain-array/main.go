package main

import "fmt"

func validMountainArray(arr []int) bool {
	if arr == nil || len(arr) < 3 {
		return false
	}
	N := len(arr)
	i := 0

	//walking up
	for i+1 < N && arr[i] < arr[i+1] {
		i++
	}

	//peak can't be first or last
	if i == 0 || i == N-1 {
		return false
	}

	//walk down
	for i+1 < N && arr[i] > arr[i+1] {
		i++
	}

	return i == N-1
}

func main() {
	input := []int{0, 1, 2, 3, 4, 8, 9, 10, 12, 12, 11}
	fmt.Println(validMountainArray(input))

}
