// https://leetcode.com/problems/binary-search/
package easy

func Search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, l int, r int, target int) int {
	for r >= l {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		}
		if nums[m] > target {
			r = m - 1
		}
		if nums[m] < target {
			l = m + 1
		}
	}
	return -1
}
