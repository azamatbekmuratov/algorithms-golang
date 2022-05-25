package main

func reverseString(s []byte) {
	helper(s, 0, len(s)-1)
}

func helper(s []byte, left int, right int) {
	if left >= right {
		return
	}

	s[left], s[right] = s[right], s[left]

	left++
	right--
	helper(s, left, right)

}
func main() {
	input := "hello"

	reverseString([]byte(input))
}
