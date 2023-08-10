package leetcode

/*

Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.

lets try and implement maps in this problem
*/

func findDisappearedNumbers(nums []int) []int {
	// struct{}{} is 0 bytes and this map is go's version of sets
	ansMap := map[int]struct{}{}
	max := len(nums)
	ans := []int{}

	for _, n := range nums {
		ansMap[n] = struct{}{}
	}

	for i := 1; i <= max; i++ {
		if _, ok := ansMap[i]; !ok {
			ans = append(ans, i)
		}
	}

	return ans
}
