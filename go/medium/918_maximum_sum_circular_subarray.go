// https://leetcode.com/problems/maximum-sum-circular-subarray/
package medium

func MaxSubarraySumCircular(nums []int) int {
	currMax := 0
	currMin := 0
	maxSum := nums[0]
	minSum := nums[0]
	total := 0

	for i := range nums {
		currMax = max(nums[i], currMax+nums[i])
		maxSum = max(maxSum, currMax)

		currMin = min(nums[i], currMin+nums[i])
		minSum = min(minSum, currMin)

		total += nums[i]
	}

	if total == minSum {
		return maxSum
	}

	return max(maxSum, total-minSum)
}
