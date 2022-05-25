package main

import "fmt"

type Stack []int

func (s *Stack) push(input int) {
	*s = append(*s, input)
}

func (s *Stack) pop() int {
	var n int = len(*s) - 1 //Top element
	var arr []int = *s
	*s = arr[:n] // Pop
	fmt.Printf("pop value: %v\n", arr[n])
	return arr[n]
}

func main() {
	stack := Stack{}

	stack.push(1)
	stack.push(2)

	stack.pop()

}
