package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNRollsTargetSum(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		testName  string
		numDice   int
		diceSides int
		target    int
		ans       int
	}{
		{
			"basic",
			1,
			6,
			3,
			1,
		},
		{
			"basic2",
			2,
			6,
			7,
			6,
		},
		{
			"basic3",
			3,
			6,
			13,
			21,
		},

		// The answer must be returned modulo 10^9 + 7
		{
			"large number",
			30,
			30,
			500,
			222616187,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			res := nRollsTargetSum(tc.numDice, tc.diceSides, tc.target)
			assert.Equal(t, tc.ans, res)
		})
	}
}
