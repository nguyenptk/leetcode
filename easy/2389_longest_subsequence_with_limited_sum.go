// https://leetcode.com/problems/longest-subsequence-with-limited-sum/
package easy

import "sort"

func AnswerQueries(nums []int, queries []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	result := []int{}
	for _, v := range queries {
		count := 0
		for j := 0; j < len(nums); j++ {
			if v-nums[j] >= 0 {
				v -= nums[j]
				count++
			} else {
				break
			}
		}
		result = append(result, count)
	}
	return result
}
