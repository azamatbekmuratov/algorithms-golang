package main

import "fmt"

func main() {
	var n = []int{84, 72, 72, 72, 84, 72}

	var isDone = false

	for !isDone {
		isDone = true
		var i = 0
		for i < len(n)-1 {
			if n[i] > n[i+1] {
				n[i], n[i+1] = n[i+1], n[i]
			}
			i++
		}
	}
	fmt.Println(n)
}
