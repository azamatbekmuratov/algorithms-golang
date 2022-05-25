package main

import "fmt"

type Rooms []int

func rob(nums []int) {
	memo := make(map[int]int)
	dp(len(nums)-1, memo)
}

func dp(i int, memo map[int]int) {
	//Base cases
	v, found := memo[3]
	if found {
		fmt.Println(v)
	}

	if v, found := memo[3]; found {
		fmt.Println(v)
	}

	if v, found := memo[3]; found {
		fmt.Println(v)
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 7}
	rob(nums)
}
