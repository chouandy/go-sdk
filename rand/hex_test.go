package rand

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHex(t *testing.T) {
	// Set test cases
	testCases := []struct {
		size     int
		expected int
	}{
		{
			size:     7,
			expected: 7,
		},
		{
			size:     10,
			expected: 10,
		},
		{
			size:     32,
			expected: 32,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			h := Hex(testCase.size)
			assert.Equal(t, testCase.expected, len(h))
		})
	}
}

func TestHexString(t *testing.T) {
	// Set test cases
	testCases := []struct {
		size     int
		expected int
	}{
		{
			size:     7,
			expected: 16,
		},
		{
			size:     10,
			expected: 22,
		},
		{
			size:     32,
			expected: 66,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			s := HexString(testCase.size)
			assert.Equal(t, testCase.expected, len(s))
		})
	}
}
