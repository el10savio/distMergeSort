package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSort ...
func TestSort(t *testing.T) {
	for _, testCase := range testSortTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			actualList := Sort(testCase.list)
			assert.Equal(t, testCase.expectedList, actualList)
		})
	}
}

// TestMerge ...
func TestMerge(t *testing.T) {
	for _, testCase := range testMergeTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			actualList := merge(testCase.list1, testCase.list2)
			assert.Equal(t, testCase.expectedList, actualList)
		})
	}
}
