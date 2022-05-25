package main

import "fmt"

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] < b[i] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	for i < len(a) {
		final = append(final, a[i])
		i++
	}

	for j < len(b) {
		final = append(final, b[j])
		j++
	}

	return final
}

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func main() {
	unsorted := []int{84, 72, 72, 72, 84, 72}
	sorted := mergeSort(unsorted)

	fmt.Println(sorted)
}
