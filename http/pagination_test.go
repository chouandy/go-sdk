package http

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPagination(t *testing.T) {
	type expected struct {
		limit  int64
		offset int64
		pages  int64
	}

	// Set test cases
	testCases := []struct {
		pagination Pagination
		expected   expected
	}{
		{
			Pagination{
				Page:      1,
				PageSize:  10,
				TotalSize: 10,
			},
			expected{
				limit:  10,
				offset: 0,
				pages:  1,
			},
		},
		{
			Pagination{
				Page:      1,
				PageSize:  10,
				TotalSize: 12,
			},
			expected{
				limit:  10,
				offset: 0,
				pages:  2,
			},
		},
		{
			Pagination{
				Page:      2,
				PageSize:  10,
				TotalSize: 12,
			},
			expected{
				limit:  10,
				offset: 10,
				pages:  2,
			},
		},
		{
			Pagination{
				Page:      0,
				PageSize:  0,
				TotalSize: 32,
			},
			expected{
				limit:  10,
				offset: 0,
				pages:  4,
			},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.expected.limit, testCase.pagination.Limit())
			assert.Equal(t, testCase.expected.offset, testCase.pagination.Offset())
			assert.Equal(t, testCase.expected.pages, testCase.pagination.Pages())
		})
	}
}
