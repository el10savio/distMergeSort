package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Move Test Suite
// To Different File

var testSortTestSuite = []struct {
	name         string
	list         []int
	expectedList []int
}{
	{"BasicFuntionality", []int{5, 7, 3, 1}, []int{1, 3, 5, 7}},
	{"EmptyList", []int{}, []int{}},
}

// TestSort ...
func TestSort(t *testing.T) {
	for _, testCase := range testSortTestSuite {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actualList := Sort(testCase.list)
			assert.Equal(t, testCase.expectedList, actualList)
		})
	}
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

// TestMerge ...
func TestMerge(t *testing.T) {
	for _, testCase := range testMergeTestSuite {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actualList := merge(testCase.list1, testCase.list2)
			assert.Equal(t, testCase.expectedList, actualList)
		})
	}
}
