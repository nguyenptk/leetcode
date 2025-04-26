// https://leetcode.com/problems/maximum-absolute-sum-of-any-subarray/
package medium

func MaxAbsoluteSum(nums []int) int {
	curr := 0
	minPrefix := 0
	maxPrefix := 0

	for _, num := range nums {
		curr += num
		minPrefix = min(minPrefix, curr)
		maxPrefix = max(maxPrefix, curr)
	}

	return maxPrefix - minPrefix
}
