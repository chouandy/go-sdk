package redis_test

import (
	. "github.com/chouandy/go-sdk/redis"

	"errors"
	"fmt"
	"testing"

	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
)

func TestIsErrNil(t *testing.T) {
	// Set test cases
	testCases := []struct {
		err      error
		expected bool
	}{
		{
			errors.New("value not found"),
			false,
		},
		{
			redis.ErrNil,
			true,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.expected, IsErrNil(testCase.err))
		})
	}
}
