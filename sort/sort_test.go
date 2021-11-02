package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSort tests the functionality
// of the Sort() function
func TestSort(t *testing.T) {
	for _, testCase := range testSortTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			actualList, actualError := Sort(testCase.list)
			assert.NoError(t, actualError)
			assert.Equal(t, testCase.expectedList, actualList)
		})
	}
}

// TestMerge tests the functionality
// of the merge() function
func TestMerge(t *testing.T) {
	for _, testCase := range testMergeTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			actualList := merge(testCase.list1, testCase.list2)
			assert.Equal(t, testCase.expectedList, actualList)
		})
	}
}

// TestCreateChunks tests the functionality
// of the createChunks() function
func TestCreateChunks(t *testing.T) {
	for _, testCase := range testCreateChunksTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			actualList := createChunks(testCase.list, testCase.chunkSize)
			assert.ElementsMatch(t, testCase.expectedList, actualList)
		})
	}
}
