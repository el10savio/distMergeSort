package sort

var testSortTestSuite = []struct {
	name         string
	list         []int
	expectedList []int
}{
	{"BasicFuntionality", []int{5, 7, 3, 1}, []int{1, 3, 5, 7}},
	{"EmptyList", []int{}, []int{}},
}

var testMergeTestSuite = []struct {
	name         string
	list1        []int
	list2        []int
	expectedList []int
}{
	{"BasicFuntionality", []int{1, 3, 5}, []int{2, 3, 9}, []int{1, 2, 3, 3, 5, 9}},
	{"List1Empty", []int{}, []int{2, 3, 9}, []int{2, 3, 9}},
	{"List2Empty", []int{1, 3, 5}, []int{}, []int{1, 3, 5}},
	{"BothListsEmpty", []int{}, []int{}, []int{}},
}

var testCreateChunksTestSuite = []struct {
	name         string
	list         []int
	chunkSize    int
	expectedList [][]int
}{
	{"BasicFuntionality", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}},
	{"BothEmpty", []int{}, 0, [][]int{}},
	{"EmptySlice", []int{}, 10, [][]int{}},
	{"ChunkSizeZero", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, [][]int{}},
	{"ChunkSizeOdd", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9, [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}, {10}}},
	{"ChunkSizeSmall", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}}},
	{"ChunkSizeOne", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, [][]int{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}, {10}}},
	{"GreaterThankSliceChunk", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 20, [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}},
}
