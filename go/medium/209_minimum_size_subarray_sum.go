// https://leetcode.com/problems/minimum-size-subarray-sum/
package medium

import "math"

func MinSubArrayLen(target int, nums []int) int {
	result := math.MaxInt
	l := 0
	sum := 0

	for r := 0; r < len(nums); r++ {
		sum += nums[r]
		for sum >= target {
			result = min(result, r-l+1)
			sum -= nums[l]
			l++
		}
	}

	if result == math.MaxInt {
		return 0
	}
	return result
}
