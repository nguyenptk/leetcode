//https://leetcode.com/problems/shortest-unsorted-continuous-subarray/
package medium

import "math"

func FindUnsortedSubarray(nums []int) int {
	min := math.MaxInt
	max := math.MinInt

	flag := false
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			flag = true
		}
		if flag {
			min = int(math.Min(float64(min), float64(nums[i])))
		}
	}

	flag = false
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] > nums[i+1] {
			flag = true
		}
		if flag {
			max = int(math.Max(float64(max), float64(nums[i])))
		}
	}

	var l int
	for l = 0; l < len(nums); l++ {
		if min < nums[l] {
			break
		}
	}

	var r int
	for r = len(nums) - 1; r >= 0; r-- {
		if max > nums[r] {
			break
		}
	}

	if r-l < 0 {
		return 0
	}
	return r - l + 1
}
