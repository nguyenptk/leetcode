// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array
package medium

import "math"

func FindMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	l := 0
	r := len(nums) - 1

	if nums[r] > nums[0] {
		return nums[0]
	}

	for l <= r {
		m := l + (r-l)/2

		if nums[m] > nums[m+1] {
			return nums[m+1]
		} else if nums[m-1] > nums[m] {
			return nums[m]
		}

		if nums[m] > nums[0] {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return math.MaxInt16
}
