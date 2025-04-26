// https://leetcode.com/problems/maximum-average-subarray-i/
package easy

import "math"

func FindMaxAverage(nums []int, k int) float64 {
	result := math.MinInt
	l := 0
	r := 0

	sum := 0
	for r < len(nums) {
		sum += nums[r]
		if r >= k-1 {
			result = max(result, sum)
			sum -= nums[l]
			l++
		}
		r++
	}

	return float64(result) / float64(k)
}
