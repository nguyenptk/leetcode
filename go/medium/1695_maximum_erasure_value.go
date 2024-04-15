// https://leetcode.com/problems/maximum-erasure-value/
package medium

import "math"

func MaximumUniqueSubarray(nums []int) int {
	// Init
	total := 0
	best := 0
	result := make([]int, 10001)
	l := 0
	r := 0

	for ; r < len(nums); r++ {
		result[nums[r]]++
		total += nums[r]
		for result[nums[r]] > 1 {
			result[nums[l]]--
			total -= nums[l]
			l++
		}
		best = int(math.Max(float64(best), float64(total)))
	}

	return best
}
