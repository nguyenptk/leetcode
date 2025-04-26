// https://leetcode.com/problems/maximum-average-subarray-ii/
package hard

import "math"

func FindMaxAverage(nums []int, k int) float64 {
	l, r := math.MaxFloat64, math.SmallestNonzeroFloat64
	for _, num := range nums {
		l = min(l, float64(num))
		r = max(r, float64(num))
	}

	precision := 1e-5
	for r-l > precision {
		m := (r + l) / 2
		if isValid(nums, m, k) {
			l = m
		} else {
			r = m
		}
	}

	return l
}

func isValid(nums []int, m float64, k int) bool {
	sum := 0.0
	minPrefixSum := 0.0
	prefixSum := 0.0

	for i := 0; i < k; i++ {
		sum += float64(nums[i]) - m
	}

	if sum >= 0 {
		return true
	}

	for i := k; i < len(nums); i++ {
		sum += float64(nums[i]) - m
		prefixSum += float64(nums[i-k]) - m
		minPrefixSum = min(minPrefixSum, prefixSum)

		if sum > minPrefixSum {
			return true
		}
	}

	return false
}
