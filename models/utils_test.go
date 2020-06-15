package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIDsDifferences(t *testing.T) {
	// Set test cases
	testCases := []struct {
		sliceA   []uint64
		sliceB   []uint64
		expected struct {
			differenceA []uint64
			differenceB []uint64
		}
	}{
		{
			sliceA: []uint64{},
			sliceB: []uint64{},
			expected: struct {
				differenceA []uint64
				differenceB []uint64
			}{
				[]uint64{},
				[]uint64{},
			},
		},
		{
			sliceA: []uint64{1, 2, 3, 4, 5, 6},
			sliceB: []uint64{},
			expected: struct {
				differenceA []uint64
				differenceB []uint64
			}{
				[]uint64{1, 2, 3, 4, 5, 6},
				[]uint64{},
			},
		},
		{
			sliceA: []uint64{},
			sliceB: []uint64{1, 2, 3, 4, 5, 6},
			expected: struct {
				differenceA []uint64
				differenceB []uint64
			}{
				[]uint64{},
				[]uint64{1, 2, 3, 4, 5, 6},
			},
		},
		{
			sliceA: []uint64{1, 2, 3, 4, 5, 6},
			sliceB: []uint64{1, 2, 3, 4, 5, 6},
			expected: struct {
				differenceA []uint64
				differenceB []uint64
			}{
				[]uint64{},
				[]uint64{},
			},
		},
		{
			sliceA: []uint64{1, 2, 3},
			sliceB: []uint64{4, 5, 6},
			expected: struct {
				differenceA []uint64
				differenceB []uint64
			}{
				[]uint64{1, 2, 3},
				[]uint64{4, 5, 6},
			},
		},
		{
			sliceA: []uint64{1, 2, 3, 4},
			sliceB: []uint64{3, 4, 5, 6},
			expected: struct {
				differenceA []uint64
				differenceB []uint64
			}{
				[]uint64{1, 2},
				[]uint64{5, 6},
			},
		},
		{
			sliceA: []uint64{1, 2, 3, 4, 5},
			sliceB: []uint64{2, 3, 4, 5, 6},
			expected: struct {
				differenceA []uint64
				differenceB []uint64
			}{
				[]uint64{1},
				[]uint64{6},
			},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			differenceA, differenceB := IDsDifferences(testCase.sliceA, testCase.sliceB)
			assert.Equal(t, testCase.expected.differenceA, differenceA)
			assert.Equal(t, testCase.expected.differenceB, differenceB)
		})
	}
}
