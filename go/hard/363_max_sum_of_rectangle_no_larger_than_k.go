// https://leetcode.com/problems/max-sum-of-rectangle-no-larger-than-k/
package hard

import (
	"math"
	"sort"
)

func MaxSumSubmatrix(matrix [][]int, k int) int {
	rows, cols := len(matrix), len(matrix[0])
	maxSum := math.MinInt32

	for left := 0; left < cols; left++ {
		// Initialize sums for rows
		rowSums := make([]int, rows)

		for right := left; right < cols; right++ {
			// Update row sums for the current right column
			for i := 0; i < rows; i++ {
				rowSums[i] += matrix[i][right]
			}

			// Now find the max subarray no larger than k
			maxSum = max(maxSum, maxSumNoLargerThanK(rowSums, k))
		}
	}

	return maxSum
}

func maxSumNoLargerThanK(nums []int, k int) int {
	maxSum := math.MinInt32
	prefixSums := []int{0} // Sorted prefix sums
	currentSum := 0

	for _, num := range nums {
		currentSum += num

		// Use binary search to find the smallest prefix sum so that currentSum - prefixSum <= k
		target := currentSum - k
		idx := sort.SearchInts(prefixSums, target)

		if idx < len(prefixSums) {
			maxSum = max(maxSum, currentSum-prefixSums[idx])
		}

		// Insert the currentSum in a sorted manner
		prefixSums = insertSorted(prefixSums, currentSum)
	}

	return maxSum
}

// Binary search insertion to maintain sorted prefixSums
func insertSorted(nums []int, target int) []int {
	idx := sort.SearchInts(nums, target)
	nums = append(nums, 0)
	copy(nums[idx+1:], nums[idx:])
	nums[idx] = target
	return nums
}
