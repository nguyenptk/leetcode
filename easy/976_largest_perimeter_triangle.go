// https://leetcode.com/problems/largest-perimeter-triangle/
package easy

import "sort"

func LargestPerimeter(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := len(nums) - 1; i > 1; i-- {
		if nums[i-2]+nums[i-1] > nums[i] {
			return nums[i-2] + nums[i-1] + nums[i]
		}
	}
	return 0
}
