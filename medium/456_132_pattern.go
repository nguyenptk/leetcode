// https://leetcode.com/problems/132-pattern/
package medium

import "math"

func Find132pattern(nums []int) bool {
	// init stack
	stack := []int{}
	second := math.MinInt // Init the second element
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < second { // In this case, nums[i] is the first element
			return true
		}
		for len(stack) > 0 && nums[i] > stack[len(stack)-1] {
			second = stack[len(stack)-1] // Override second by Top stack
			stack = stack[:len(stack)-1] // Pop stack
		}
		stack = append(stack, nums[i]) // Push nums[i] to stack
	}
	return false
}
