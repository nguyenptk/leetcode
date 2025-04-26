// https://leetcode.com/problems/kth-largest-element-in-an-array/
package medium

import "github.com/emirpasic/gods/queues/priorityqueue"

func FindKthLargest(nums []int, k int) int {
	queue := priorityqueue.NewWith(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})

	for _, num := range nums {
		queue.Enqueue(num)
		if queue.Size() > k {
			queue.Dequeue()
		}
	}

	val, _ := queue.Dequeue()
	return val.(int)
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
