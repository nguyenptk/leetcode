// https://leetcode.com/problems/maximum-subarray/
package easy

import "math"

func MaxSubArray(nums []int) int {
	// Divide and Conquer algorithm
	// n := len(nums)
	// return maxSubArraySum(nums, 0, n-1)

	// Kadane's Algorithm
	return kadaneAlgorithm(nums)
}

// Return sum of maximum sub array from left to right
func maxSubArraySum(nums []int, l int, r int) int {

	// length is 1 -> return
	if l == r {
		return nums[l]
	}

	// get mid position of array
	m := (l + r) / 2

	// get max of left to mid array
	maxLeft := float64(maxSubArraySum(nums, l, m))
	// get max of mid to right array
	maxRight := float64(maxSubArraySum(nums, m+1, r))
	// find max of left and right array
	maxSubArray := math.Max(maxLeft, maxRight)

	// get max of crossing array
	maxCrossing := float64(maxCrossingSum(nums, l, m, r))

	// find max of full array & crossing array
	return int(math.Max(maxSubArray, maxCrossing))
}

// Find the maximum sum in array
func maxCrossingSum(nums []int, l int, m int, r int) int {

	// find max of left array
	sum := 0
	leftSum := math.MinInt32
	for i := m; i > l; i-- {
		sum = sum + nums[i]
		if sum > leftSum {
			leftSum = sum
		}
	}

	// find max of right array
	sum = 0
	rightSum := math.MinInt32
	for i := m + 1; i <= r; i++ {
		sum = sum + nums[i]
		if sum > rightSum {
			rightSum = sum
		}
	}

	// sum max of left and right array
	sumLeftRight := float64(leftSum + rightSum)

	// find max of left and right array
	maxLeftRight := math.Max(float64(leftSum), float64(rightSum))

	// compare sumLeftRight & maxLeftRight
	return int(math.Max(sumLeftRight, maxLeftRight))
}

func kadaneAlgorithm(nums []int) int {
	n := len(nums)
	maxSoFar := math.MinInt
	maxEndingHere := 0
	for i := 0; i < n; i++ {
		maxEndingHere = maxEndingHere + nums[i]
		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}
		if maxEndingHere < 0 {
			maxEndingHere = 0
		}
	}
	return maxSoFar
}
