// https://leetcode.com/problems/minimum-operations-to-exceed-threshold-value-ii/
package medium

import "github.com/emirpasic/gods/queues/priorityqueue"

func MinOperationsThreshold(nums []int, k int) int {
	result := 0

	minHeap := priorityqueue.NewWith(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})

	for i := 0; i < len(nums); i++ {
		minHeap.Enqueue(nums[i])
	}

	minX, _ := minHeap.Peek()
	if minX.(int) >= k {
		return 0
	}

	for minHeap.Size() > 1 {
		min1, _ := minHeap.Dequeue()
		min2, _ := minHeap.Dequeue()
		newNum := min(min1.(int), min2.(int))*2 + max(min1.(int), min2.(int))
		minHeap.Enqueue(newNum)
		result++

		minX, _ := minHeap.Peek()
		if minX.(int) >= k {
			return result
		}
	}

	return result
}
