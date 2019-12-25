package log

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestToLogrusFields(t *testing.T) {
	// New struct
	type v struct {
		Field1 string `json:"field_1"`
		Field2 string `json:"field_2"`
	}

	// Set test cases
	testCases := []struct {
		v        v
		expected logrus.Fields
	}{
		{
			v: v{
				Field1: "a",
				Field2: "b",
			},
			expected: logrus.Fields{
				"field_1": "a",
				"field_2": "b",
			},
		},
		{
			v: v{
				Field1: "c",
				Field2: "d",
			},
			expected: logrus.Fields{
				"field_1": "c",
				"field_2": "d",
			},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.expected, ToLogrusFields(testCase.v))
		})
	}
}
