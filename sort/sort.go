package sort

import "sort"

// Sort ...
func Sort(list []int) ([]int, error) {
	if len(list) == 0 {
		return []int{}, nil
	}

	if len(list) >= 10 {
		return peerSort(list)
	}

	sliceSort(list)
	return list, nil
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
