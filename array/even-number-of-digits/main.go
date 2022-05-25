//Given an array nums of integers, return how many of them contain an even number of digits.
package main

import "fmt"

func countDigits(input int) int {
	count := 1

	for input != 0 {
		input = (input / 10)
		if input > 0 {
			count++
		}
	}
	return count
}

func main() {
	input := []int{555, 901, 482, 1771}
	countEven := 0
	for _, v := range input {
		item := countDigits(v)
		if item%2 == 0 {
			countEven++
		}
	}

	fmt.Println(countEven)

}
