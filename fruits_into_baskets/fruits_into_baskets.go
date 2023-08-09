package leetcode

/*
You are visiting a farm that has a single row of fruit trees arranged from left to right. The trees are represented by an integer array fruits where fruits[i] is the type of fruit the ith tree produces.

You want to collect as much fruit as possible. However, the owner has some strict rules that you must follow:

    - You only have two baskets, and each basket can only hold a single type of fruit. There is no limit on the amount of fruit each basket can hold.
    - Starting from any tree of your choice, you must pick exactly one fruit from every tree (including the start tree) while moving to the right. The picked fruits must fit in one of your baskets.
    - Once you reach a tree with fruit that cannot fit in your baskets, you must stop.

Given the integer array fruits, return the maximum number of fruits you can pick.


tbh i thought this was longest subsequence problem, but i think using a map is the best way to approach this problem

1. iterating through the fruit trees we can add it our map (basket in the context of this problem)
2. if at any point the there are more than 2 different maps
*/

func totalFruit(fruits []int) int {
	collectedFruits := map[int]int{}
	start := 0

	for _, f := range fruits {
		if _, ok := collectedFruits[f]; !ok {
			collectedFruits[f] = 1
		} else {
			collectedFruits[f] += 1
		}

		if len(collectedFruits) > 2 {
			collectedFruits[fruits[start]] -= 1
			if collectedFruits[fruits[start]] == 0 {
				delete(collectedFruits, fruits[start])
			}
			start += 1
		}
	}
	return len(fruits) - start
}
