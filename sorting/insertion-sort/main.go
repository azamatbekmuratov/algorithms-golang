package main

import "fmt"

func main() {
	var n = []int{33, 1, 23, 49, 22, 89, 2}

	var i = 1
	for i < len(n) {
		var j = i
		for j >= 1 && n[j] < n[j-1] {
			n[j], n[j-1] = n[j-1], n[j]
			j--
		}
		i++
	}

	fmt.Println(n)
}
