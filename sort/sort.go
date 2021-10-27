package sort

import "sort"

// Sort ...
func Sort(elements []int) []int {
	if len(elements) == 0 {
		return []int{}
	}
	return sliceSort(elements)
}

// sliceSort ...
func sliceSort(elements []int) []int {
	return sort.IntSlice(elements)
}
