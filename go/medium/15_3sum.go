// https://leetcode.com/problems/3sum
package medium

import "sort"

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)

	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i-1] != nums[i] {
			l, r := i+1, len(nums)-1
			for l < r {
				sum := nums[i] + nums[l] + nums[r]
				if sum < 0 {
					l++
				} else if sum > 0 {
					r--
				} else {
					result = append(result, []int{nums[i], nums[l], nums[r]})
					l++
					r--
					for l < r && nums[l] == nums[l-1] {
						l++
					}
				}
			}
		}
	}

	return result
}
