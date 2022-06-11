// https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/
package medium

import "math"

func MinOperations(nums []int, x int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	remainder := total - x
	n := len(nums)
	maxLen := -1
	current := 0

	j := 0
	for i := 0; i < n; i++ {
		current += nums[i]
		for current > remainder && j <= i {
			current -= nums[j]
			j++
		}
		if current == remainder {
			maxLen = int(math.Max(float64(maxLen), float64(i-j+1)))
		}
	}
	if maxLen == -1 {
		return maxLen
	}
	return n - maxLen
}
