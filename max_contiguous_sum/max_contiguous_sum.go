package leetcode

import "math"

// Find the maximum sum of a subarray in an array

//  The following solution is a greedy solution (Kadane's algorithm)
// 1. Keep an ongoing sum currentSum of the array,
//  		if currentSum reset it to 0
// 				bc it could the array at that point would negatively contribute to the maxSum
// 			update s (temp var) index to the current index for startIndex
// 				why not update the ACTUAL startIndex?
//					we don't know if there will be an actual new maximum
// 2. if the currentSum > maxSum
//		maxSum = currentSum
// 		endIndex = the current index
// 		why isn't startIndex updated here?
// 			startIndex represents the start of the arr while we are iterating through it
//			updating it here wouldn't make sense

// I don't think the solution is intuitive. This problem could be solved with prep or awareness
// of a solution

func maxContiguousSum(arr []int) ([]int, int) {
	s, startIndex, endIndex := 0, 0, 0
	currentSum := 0
	maxSum := int(math.Inf(-1))

	for i, num := range arr {
		currentSum += num

		if currentSum > maxSum {
			maxSum = currentSum
			startIndex = s
			endIndex = i
		}

		if currentSum < 0 {
			currentSum = 0
			s = i + 1
		}
	}

	sum := 0
	for _, num := range arr[startIndex : endIndex+1] {
		sum += num
	}

	return []int{startIndex, endIndex}, sum
}
