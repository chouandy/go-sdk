package time

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNextWeekday(t *testing.T) {

	// Set test cases
	testCases := []struct {
		date     string
		weekday  time.Weekday
		expected string
	}{
		{
			date:     "2020-06-01",
			weekday:  time.Friday,
			expected: "2020-06-05",
		},
		{
			date:     "2020-06-02",
			weekday:  time.Friday,
			expected: "2020-06-05",
		},
		{
			date:     "2020-06-01",
			weekday:  time.Saturday,
			expected: "2020-06-06",
		},
		{
			date:     "2020-04-29",
			weekday:  time.Friday,
			expected: "2020-05-01",
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			date, err := time.Parse("2006-01-02", testCase.date)
			assert.Nil(t, err)

			nextWeekday := NextWeekday(date, testCase.weekday)

			expected, err := time.Parse("2006-01-02", testCase.expected)
			assert.Nil(t, err)
			assert.Equal(t, expected, nextWeekday)
		})
	}
}
