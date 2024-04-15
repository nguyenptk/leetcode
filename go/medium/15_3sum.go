// https://leetcode.com/problems/3sum
package medium

import "sort"

func ThreeSum(nums []int) [][]int {
	var result *[][]int

	// sort the array
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > 0 {
			// Remaining values cannot sum to 0
			break
		}
		if i == 0 || nums[i-1] != nums[i] {
			twoSum(nums, i, result)
		}
	}

	return *result
}

func twoSum(nums []int, i int, result *[][]int) {
	// init 2 pointers
	l := 1 // we start at the i = 0, so we need to init l = 1
	r := len(nums) - 1

	for l < r {
		if nums[i]+nums[l]+nums[r] < 0 {
			l++
		}
		if nums[i]+nums[l]+nums[r] > 0 {
			r--
		} else {
			*result = append(*result, []int{nums[i], nums[l], nums[r]})
			l++
			r--
			for l < r && nums[l] == nums[l-1] {
				l++
			}
		}
	}
}
