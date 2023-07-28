package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxContiguousSum(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		testName string
		day      int
		month    int
		year     int
		ans      string
	}{
		{
			"basic",
			31,
			8,
			2019,
			"Saturday",
		},
		{
			"basic2",
			18,
			7,
			1999,
			"Sunday",
		},
		{
			"basic3",
			15,
			8,
			1993,
			"Sunday",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			res := dayOfTheWeek(tc.day, tc.month, tc.year)
			assert.Equal(t, tc.ans, res)
		})
	}
}
