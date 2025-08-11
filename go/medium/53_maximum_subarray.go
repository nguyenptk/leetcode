// https://leetcode.com/problems/maximum-subarray/
package medium

func MaxSubArray(nums []int) int {
	result := nums[0]
	currSum := nums[0]
	for i := 1; i < len(nums); i++ {
		currSum = max(nums[i], currSum+nums[i])
		result = max(result, currSum)
	}
	return result
}
