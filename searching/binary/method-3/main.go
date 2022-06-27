package main

func binarySearch(a []int, search int) (result int, searchCount int) {
	mid := len(a) / 2
	switch {
	case len(a) == 0:
		result = -1 // result not found
	case a[mid] > search:
		result, searchCount = binarySearch(a[:mid], search)
	case a[mid] < search:
		result, searchCount = binarySearch(a[mid:], search)
		if result >= 0 { // if anything but the -1 "not found" result
			result += mid + 1
		}
	default:
		result = mid // found
	}
	searchCount++
	return
}

func main() {

}
