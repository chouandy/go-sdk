package validator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckPassword(t *testing.T) {
	// Set test cases
	testCases := []struct {
		password string
		n        int
		number   bool
		lower    bool
		upper    bool
		special  bool
		expected bool
	}{
		{
			password: "0Aa~",
			n:        4,
			number:   true,
			lower:    true,
			upper:    true,
			special:  true,
			expected: true,
		},
		{
			password: "Aa~",
			n:        3,
			number:   true,
			lower:    true,
			upper:    true,
			special:  true,
			expected: false,
		},
		{
			password: "Aa~",
			n:        3,
			number:   false,
			lower:    true,
			upper:    true,
			special:  true,
			expected: true,
		},
		{
			password: "0a~",
			n:        3,
			number:   true,
			lower:    true,
			upper:    true,
			special:  true,
			expected: false,
		},
		{
			password: "0a~",
			n:        3,
			number:   true,
			lower:    false,
			upper:    true,
			special:  true,
			expected: true,
		},
		{
			password: "0A~",
			n:        3,
			number:   true,
			lower:    true,
			upper:    true,
			special:  true,
			expected: false,
		},
		{
			password: "0A~",
			n:        3,
			number:   true,
			lower:    true,
			upper:    false,
			special:  true,
			expected: true,
		},
		{
			password: "0Aa",
			n:        3,
			number:   true,
			lower:    true,
			upper:    true,
			special:  true,
			expected: false,
		},
		{
			password: "0Aa",
			n:        3,
			number:   true,
			lower:    true,
			upper:    true,
			special:  false,
			expected: true,
		},
		{
			password: "0Aa",
			n:        4,
			number:   true,
			lower:    true,
			upper:    true,
			special:  false,
			expected: false,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.expected, CheckPassword(
				testCase.password,
				testCase.n,
				testCase.number,
				testCase.upper,
				testCase.lower,
				testCase.special,
			))
		})
	}
}
