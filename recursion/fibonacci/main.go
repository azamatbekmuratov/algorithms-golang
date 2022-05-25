package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}

	result := fib(n-1) + fib(n-2)
	return result
}

func main() {
	input := 3
	fmt.Println(fib(input))
}
