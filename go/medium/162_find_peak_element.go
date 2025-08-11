// https://leetcode.com/problems/find-peak-element/
package medium

func FindPeakElement(nums []int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		m := l + (r-l)/2
		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
