package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxContiguousSum(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		testName string
		fruits   []int
		ans      int
	}{
		{
			"basic",
			[]int{1, 2, 1},
			3,
		},
		{
			"basic2",
			[]int{0, 1, 2, 2},
			3,
		},
		{
			"basic3",
			[]int{1, 2, 3, 2, 2},
			4,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			res := totalFruit(tc.fruits)
			assert.Equal(t, tc.ans, res)
		})
	}
}
