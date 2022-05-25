package main

import "fmt"

func removeElement(nums []int, val int) int {
	cnt := 0
	p1 := 0
	p2 := len(nums) - 1
	if len(nums) == 1 && nums[0] == val {
		nums = nil
		return 0
	}
	for p1 < p2 {
		if nums[p1] == val {
			if nums[p2] != val {
				nums[p1] = nums[p2]
				nums[p2] = 99
				p2--
			} else {
				for p2 > p1 {
					if nums[p2] == val {
						nums[p2] = 99
						p2--
					} else {
						nums[p1] = nums[p2]
						nums[p2] = 99
						p2--
						break
					}
				}
			}

		}

		p1++
	}
	for _, v := range nums {
		if v != 99 {
			cnt++
		}
	}
	fmt.Println(nums)
	return cnt
}

func main() {
	var input []int = []int{0, 1, 2, 2, 3, 0, 4, 2}

	removeElement(input, 2)
}
