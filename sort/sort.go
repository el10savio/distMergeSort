package sort

import "sort"

// Sort is the handler that decides based on the size
// of the list wheteher to use a distributed merge sort
// across the cluster of just sort it in
// memory on the host node
func Sort(list []int) ([]int, error) {
	const largeSortCount = 100

	if len(list) == 0 {
		return []int{}, nil
	}

	if len(list) >= largeSortCount {
		return peerSort(list)
	}

	sliceSort(list)
	return list, nil
}

// sliceSort uses the standard library sort
// method to sort the list in place
func sliceSort(list []int) {
	sort.Ints(list)
}

// merge takes in two sorted lists and recursively
// merges them together to generate a
// single merged sorted list
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
