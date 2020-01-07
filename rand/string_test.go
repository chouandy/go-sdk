package rand

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	// Set test cases
	testCases := []struct {
		n        int
		expected int
	}{
		{
			n:        7,
			expected: 7,
		},
		{
			n:        10,
			expected: 10,
		},
		{
			n:        15,
			expected: 15,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			s := String(Alphanumeric, testCase.n)
			assert.Equal(t, testCase.expected, len(s))
		})
	}
}
