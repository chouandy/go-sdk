package rand

import (
	"fmt"
	"testing"

	"github.com/chouandy/go-sdk/validator"
	"github.com/stretchr/testify/assert"
)

func TestPassword(t *testing.T) {
	// Set test cases
	testCases := []struct {
		n int
	}{
		{
			n: 8,
		},
		{
			n: 12,
		},
		{
			n: 16,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			password := Password(testCase.n, true, true, true, true)
			assert.True(t, validator.CheckPassword(password, testCase.n, true, true, true, true))
		})
	}
}
