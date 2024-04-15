// https://leetcode.com/problems/subsets-ii
package medium

import "sort"

func SubsetsWithDup(nums []int) [][]int {
	result := make([][]int, 0)
	curr := make([]int, 0)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var backtrack func(idx int)
	backtrack = func(idx int) {
		result = append(result, append([]int{}, curr...))
		for i := idx; i < len(nums); i++ {
			if i > idx && nums[i] == nums[i-1] {
				continue
			}
			curr = append(curr, nums[i])
			backtrack(i + 1)
			curr = curr[:len(curr)-1]
		}
	}

	backtrack(0)
	return result
}
