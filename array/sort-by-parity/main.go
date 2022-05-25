package main

func sortArrayByParity(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	p1 := 0
	p2 := 1

	for j := p2; p2 < len(nums); j++ {
		if nums[p1]%2 != 0 {
			if nums[p2]%2 != 0 {
				p2++
			} else {
				tmp := nums[p1]
				nums[p1] = nums[p2]
				nums[p2] = tmp
				p1++
				p2++
			}
		} else {
			p1++
			p2++
		}
	}
	return nums
}

func main() {
	input := []int{0, 2}
	sortArrayByParity(input)
}
