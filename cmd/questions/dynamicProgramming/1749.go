// Maximum absolute sum of any array
// https://leetcode.com/problems/maximum-absolute-sum-of-any-subarray/description/

package main

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func maxAbsoluteSum(nums []int) int {
	maxSum, maxPositive, maxNegative := 0, 0, 0
	for _, num := range nums {
		maxPositive = max(maxPositive, 0) + num
		maxNegative = min(maxNegative, 0) + num

		currSum := max(maxPositive, -maxNegative)
		maxSum = max(maxSum, currSum)
	}

	return maxSum
}
