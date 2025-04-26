// https://leetcode.com/problems/valid-triangle-number/
package medium

import "slices"

func triangleNumber(nums []int) int {
	result := 0

	slices.Sort(nums)
	for i := 1; i < len(nums); i++ {
		l := 0
		r := i - 1
		for l < r {
			if nums[l]+nums[r] > nums[i] {
				result += r - l
				r--
			} else {
				l++
			}
		}
	}

	return result
}
