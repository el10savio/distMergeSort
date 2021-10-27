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

// merge ...
func merge(elements1, elements2 []int) []int {
	if len(elements1) == 0 {
		return elements2
	}

	if len(elements2) == 0 {
		return elements1
	}

	if elements1[0] < elements2[0] {
		return append([]int{elements1[0]}, merge(elements1[1:], elements2)...)
	}

	return append([]int{elements2[0]}, merge(elements1, elements2[1:])...)
}
