// https://leetcode.com/problems/product-of-array-except-self
package medium

func ProductExceptSelf(nums []int) []int {
	// nums = [1,2,3,4]
	result := make([]int, len(nums))
	prefix := 1

	// from left to right
	for i := 0; i < len(nums); i++ {
		// result = [1,0,0,0]
		// result = [1,1,0,0]
		// result = [1,1,2,0]
		// result = [1,1,2,6]
		result[i] = prefix
		prefix *= nums[i]
	}

	// from right to left
	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		// result = [1,1,2,6]
		// result = [1,1,8,6]
		// result = [1,12,8,6]
		// result = [24,12,8,6]
		result[i] *= postfix
		postfix *= nums[i]
	}

	return result
}
