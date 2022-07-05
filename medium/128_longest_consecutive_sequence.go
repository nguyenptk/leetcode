// https://leetcode.com/problems/longest-consecutive-sequence/
package medium

import "sort"

func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Sort the array of numbers
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	// DP
	result := 0
	count := 0

	for i := 0; i < len(nums)-1; i++ {
		if nums[i]+1 == nums[i+1] { // Check consecutive sequence
			count++
		} else if nums[i] == nums[i+1] { // Check duplicated
			continue
		} else {
			count = 0 // Reset counter
		}
		if count > result {
			result = count // Compare and override result by counter
		}
	}

	return result + 1
}
