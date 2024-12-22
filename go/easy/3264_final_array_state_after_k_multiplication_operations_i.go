// https://leetcode.com/problems/final-array-state-after-k-multiplication-operations-i/
package easy

import "github.com/emirpasic/gods/queues/priorityqueue"

func GetFinalState(nums []int, k int, multiplier int) []int {
	n := len(nums)

	minHeap := priorityqueue.NewWith(func(a, b interface{}) int {
		aVal, aIdx := a.([]int)[0], a.([]int)[1]
		bVal, bIdx := b.([]int)[0], b.([]int)[1]
		if aVal == bVal {
			return aIdx - bIdx
		}
		return aVal - bVal
	})
	for i := 0; i < n; i++ {
		minHeap.Enqueue([]int{nums[i], i})
	}

	for ; k > 0; k-- {
		// Dequeue
		smallest, _ := minHeap.Dequeue()
		val := smallest.([]int)[0]
		idx := smallest.([]int)[1]

		// Update value
		val *= multiplier
		nums[idx] = val

		// Enqueue new value
		minHeap.Enqueue([]int{val, idx})
	}

	return nums
}
