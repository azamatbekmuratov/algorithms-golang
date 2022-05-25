package main

import "fmt"

func checkIfExist(arr []int) bool {
	seen := make(map[int]bool)

	for i := 0; i < len(arr); i++ {
		_, ok := seen[2*arr[i]]
		if ok {
			return true
		}
		if arr[i]%2 == 0 {
			_, ok := seen[arr[i]/2]
			if ok {
				return true
			}
		}
		seen[arr[i]] = true
	}
	return false
}

func main() {
	input := []int{-10, 0, -3, 3}
	fmt.Println(checkIfExist(input))
}
