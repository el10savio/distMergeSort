package sort

import "sort"

// Sort ...
func Sort(list []int) []int {
	if len(list) == 0 {
		return []int{}
	}

	sliceSort(list)

	return list
}

// sliceSort ...
func sliceSort(list []int) {
	sort.Ints(list)
}

// merge ...
func merge(list1, list2 []int) []int {
	if len(list1) == 0 {
		return list2
	}

	if len(list2) == 0 {
		return list1
	}

	if list1[0] < list2[0] {
		return append([]int{list1[0]}, merge(list1[1:], list2)...)
	}

	return append([]int{list2[0]}, merge(list1, list2[1:])...)
}
