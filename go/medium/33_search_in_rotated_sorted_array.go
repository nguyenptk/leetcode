// https://leetcode.com/problems/search-in-rotated-sorted-array
package medium

func Search(nums []int, target int) int {
	n := len(nums)
	l := 0
	r := n - 1
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		} else if nums[m] >= nums[l] {
			// Left is sorted
			// 0,1,2,3,4,5,6,7
			// 1,2,3,4,5,6,7,0
			// 2,3,4,5,6,7,0,1
			// 3,4,5,6,7,0,1,2
			if target >= nums[l] && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			// Right is sorted
			// 4,5,6,7,0,1,2,3
			// 5,6,7,0,1,2,3,4
			// 6,7,0,1,2,3,4,5
			// 7,0,1,2,3,4,5,6
			if target <= nums[r] && target > nums[m] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}

	return -1
}
