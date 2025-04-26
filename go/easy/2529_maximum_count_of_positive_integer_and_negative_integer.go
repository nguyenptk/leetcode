// https://leetcode.com/problems/maximum-count-of-positive-integer-and-negative-integer/
package easy

func MaximumCount(nums []int) int {
	pos := len(nums) - upperBound(nums)
	neg := lowerBound(nums)
	return max(pos, neg)
}

func lowerBound(nums []int) int {
	idx := len(nums)
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] < 0 {
			l = m + 1
		} else {
			r = m - 1
			idx = m
		}
	}
	return idx
}

func upperBound(nums []int) int {
	idx := len(nums)
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] <= 0 {
			l = m + 1
		} else {
			r = m - 1
			idx = m
		}
	}
	return idx
}
