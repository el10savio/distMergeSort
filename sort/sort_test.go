package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
