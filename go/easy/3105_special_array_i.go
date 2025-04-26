// https://leetcode.com/problems/longest-strictly-increasing-or-strictly-decreasing-subarray/
package easy

func LongestMonotonicSubarray(nums []int) int {
	des := 1
	inc := 1
	maxLen := 1

	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] {
			inc++
			des = 1
		} else if nums[i+1] < nums[i] {
			des++
			inc = 1
		} else {
			inc = 1
			des = 1
		}
		maxLen = max(maxLen, max(inc, des))
	}

	return maxLen
}
