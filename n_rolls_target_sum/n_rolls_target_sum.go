package leetcode

/*
You have n dice, and each die has k faces numbered from 1 to k.
Given three integers n, k, and target, return the number of possible ways (out of the k^n total ways) to roll the dice,
	so the sum of the face-up numbers equals target. Since the answer may be too large, return it modulo 10^9 + 7.

What is the base case?

	2 dice, 6 sides, target = 7
		if dice 1 is a, and dice 2 is 1, then a = 6
		if dice 1 is a, and dice 2 is 2, then a = 5
		if dice 1 is a, and dice 2 is 3, then a = 4
		if dice 1 is a, and dice 2 is 4, then a = 3
		if dice 1 is a, and dice 2 is 5, then a = 2
		if dice 1 is a, and dice 2 is 6, then a = 1

	3 dice, 6 sides, target = 13
		if dice 1 is 1, then the remaining dice sum to 12
		if dice 1 is 2, then the remaining dice sum to 11
		if dice 1 is 3, then the remaining dice sum to 10
		if dice 1 is 4, then the remaining dice sum to 9
		if dice 1 is 5, then the remaining dice sum to 8
		if dice 1 is 6, then the remaining dice sum to 7
	Our parameters our number of dice (d), number of sides (s), and the target (t).
	Find dp(d, s, t) when dp(3, 6, 13)
		From our (3, 6, 13)
		remaining dice sum to 12 = dp(2, 6, 12) [(tautology btw) it takes me forever to conceptually understand in notation]
		remaining dice sum to 11 = dp(2, 6, 11)
		remaining dice sum to 10 = dp(2, 6, 10)
		remaining dice sum to 9 = dp(2, 6, 9)
		remaining dice sum to 8 = dp(2, 6, 8)
		remaining dice sum to 7 = dp(2, 6, 7)

	And since there are only 6 cases, the sum of all of dp(2, 6 ,range(7,12)) = dp(3, 6, 13), sooo try find the pattern to write the recursion relation

	dp(d, s, t) = dp(d-1, s, t-1) + dp(d-1, s, t-2) + dp(d-1, s, t-3) + ... + dp(d-1, s, t-s)

	OKAY... NOW WE ARE GETTING SOMEWHERE....

	s is a constant... what is the base case?

	d = 1? if d = 0, then d-1 doesn't make sense
	s > 0? idk these seems like logical bounds rather than math-y ones

	--------------------------------
	(i had to get hints here)
	base case is dp(d, s, t) = 0 (duh) bc which is when you have 0 dice, and your target can only be 0.
		which is dp(0, s, 0) = 0

	dp(d, s, t) = number of ways you can have the combination of d dice, s sides, and t target

	dp(d, s, t) != t target (which is conceptually hard for me to wrap my head around)
	--------------------------------

	??? okay do i have enough to start coding, what does my table of solutions look like?
*/

func nRollsTargetSum(n, k, target int) int {
	memoizationCache := map[[2]int]int{}

	return dp(n, k, target, memoizationCache)
}

func dp(n, k, target int, cache map[[2]int]int) int {
	// these are the variable values according to our recurrence relations
	ans := [2]int{n, target}
	nCombinations := 0

	if val, ok := cache[ans]; ok {
		return val
	}

	// dp(d, s, t) = dp(d-1, s, t-1) + dp(d-1, s, t-2) + dp(d-1, s, t-3) + ... + dp(d-1, s, t-s)
	// WHERE DOES THE BASE CASE COME IN
	for i := 1; i <= k; i++ {
		// case when there is 1 dice
		// i think this makes sense because we don't technically need the 0 values for our table
		if target == i && n == 1 {
			nCombinations += 1
		}

		if target > i && n > 1 {
			nCombinations += dp(n-1, k, target-i, cache)
		}
	}

	nCombinations = nCombinations % 1000000007
	cache[ans] = nCombinations
	// fmt.Println(nCombinations)
	return nCombinations
}

// So I understand the solution, but putting it together myself would require a lot of time
// I need to work on dynamic programming more q.q
