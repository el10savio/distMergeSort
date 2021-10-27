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
