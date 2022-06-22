// https://leetcode.com/problems/kth-largest-element-in-an-array/
package medium

import "sort"

func FindKthLargest(nums []int, k int) int {
	n := len(nums)
	// Optimize test case with len(nums) == 1
	if n == 1 {
		return nums[0]
	}

	// Sort
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	kthLargest := nums[0] // First
	nums = nums[1:]       // Poll

	for i := 0; i < n-k; i++ {
		kthLargest = nums[0] // First
		nums = nums[1:]      // Poll
	}

	return kthLargest
}

// Solution 2
// func FindKthLargest(nums []int, k int) int {
// 	// Optimize test case with len(nums) == 1
// 	if len(nums) == 1 {
// 		return nums[0]
// 	}

// 	// Sort
// 	sort.Slice(nums, func(i, j int) bool {
// 		return nums[i] < nums[j]
// 	})

// 	return nums[len(nums)-k]
// }
