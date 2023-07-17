package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxContiguousSum(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		testName string
		arg1     []int
		ans      []int
		s        int
	}{
		{
			"basic",
			[]int{1, 2, 3, 4, 5},
			[]int{0, 4},
			15,
		},
		{
			"mixed negatives and positives",
			[]int{-2, -3, 4, -1, -2, 1, 5, -3},
			[]int{2, 6},
			7,
		},
		{
			"mixed negatives and positives 2",
			[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			[]int{3, 6},
			6,
		},
		{
			"one element",
			[]int{-3},
			[]int{0, 0},
			-3,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			res, s := maxContiguousSum(tc.arg1)
			assert.Equal(t, tc.ans, res)
			assert.Equal(t, tc.s, s)
		})
	}
}
