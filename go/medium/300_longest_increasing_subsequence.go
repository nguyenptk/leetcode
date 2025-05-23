// https://leetcode.com/problems/longest-increasing-subsequence/
package medium

import "sort"

func LengthOfLIS(nums []int) int {
	return biSearchLengthOfLIS(nums)
	// return dpLengthOfLIS(nums)
}

func biSearchLengthOfLIS(nums []int) int {
	list := []int{}
	for _, num := range nums {
		if len(list) == 0 || num > list[len(list)-1] {
			list = append(list, num)
		} else {
			j := sort.SearchInts(list, num)
			list[j] = num
		}
	}
	return len(list)
}

func dpLengthOfLIS(nums []int) int {
	n := len(nums)

	dp := make([]int, n)
	for i := range nums {
		dp[i] = 1
	}
	dp[0] = 1

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j+1])
			}
		}
	}

	longest := 0
	for i := range dp {
		longest = max(longest, dp[i])
	}

	return longest
}
