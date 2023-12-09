package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDisappearedNumbers(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		testName string
		arg1     []int
		ans      []int
	}{
		{
			"tc1",
			[]int{4, 3, 2, 7, 8, 2, 3, 1},
			[]int{5, 6},
		},
		{
			"tc2",
			[]int{1, 1},
			[]int{2},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			res := findDisappearedNumbers(tc.arg1)
			assert.Equal(t, tc.ans, res)
		})
	}
}
