// https://leetcode.com/problems/rotate-array/
package medium

// Add additional return value for unit-test
func Rotate(nums []int, k int) []int {
	// if k < len(nums)       -> len(nums) - k
	// else if k == len(nums) -> len(nums)
	// else                   -> len(nums) - (k-len(nums))
	r := len(nums) - k%len(nums)
	copy(nums, append(nums[r:], nums[:r]...))
	return nums
}
