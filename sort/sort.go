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
	// Check if either array is empty - if so return the other

	// Get bigger array size to allocate

	// Iterate over bigger array & get smaller element

	// Push smaller element to merged list

	// Redo until any array gets empty

	// Merge the other non empty array

	// Return merged array

	return []int{}
}
