package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNRollsTargetSum(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		testName    string
		romanString string
		ans         int
	}{
		{
			"I - 1",
			"I",
			1,
		},
		{
			"MCMXCIV - 1994",
			"MCMXCIV",
			1994,
		},
		{
			"LVIII - 58",
			"LVIII",
			58,
		},
		{
			"MMVIII",
			"MMVIII",
			2008,
		},
		{
			"MCMXCIII - 1993",
			"MCMXCIII",
			1993,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			res := romanStringToNum(tc.romanString)
			assert.Equal(t, tc.ans, res)
		})
	}
}
