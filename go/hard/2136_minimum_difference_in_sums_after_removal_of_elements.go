// https://leetcode.com/problems/minimum-difference-in-sums-after-removal-of-elements/
package hard

import (
	"math"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

func newMaxHeap() *priorityqueue.Queue {
	return priorityqueue.NewWith(func(a, b interface{}) int {
		return int(b.(int64) - a.(int64))
	})
}

func newMinHeap() *priorityqueue.Queue {
	return priorityqueue.NewWith(func(a, b interface{}) int {
		return int(a.(int64) - b.(int64))
	})
}

func MinimumDifference(nums []int) int64 {
	n := len(nums)
	k := n / 3

	dpLeft := make([]int64, n)
	maxHeap := newMaxHeap()
	var sumLeft int64
	for i := 0; i < k; i++ {
		sumLeft += int64(nums[i])
		maxHeap.Enqueue(int64(nums[i]))
	}
	dpLeft[k-1] = sumLeft
	for i := k; i < n; i++ {
		x := int64(nums[i])
		// peek max
		node, _ := maxHeap.Peek()
		top := node.(int64)
		if top > x {
			sumLeft += x - top
			maxHeap.Dequeue()
			maxHeap.Enqueue(x)
		}
		dpLeft[i] = sumLeft
	}

	dpRight := make([]int64, n)
	minHeap := newMinHeap()
	var sumRight int64
	for i := n - 1; i >= n-k; i-- {
		sumRight += int64(nums[i])
		minHeap.Enqueue(int64(nums[i]))
	}
	dpRight[n-k] = sumRight
	for i := n - k - 1; i >= 0; i-- {
		x := int64(nums[i])
		// peek min
		node, _ := minHeap.Peek()
		top := node.(int64)
		if top < x {
			sumRight += x - top
			minHeap.Dequeue()
			minHeap.Enqueue(x)
		}
		dpRight[i] = sumRight
	}

	result := int64(math.MaxInt64)
	for i := k - 1; i < 2*k; i++ {
		diff := dpLeft[i] - dpRight[i+1]
		if result < diff {
			result = diff
		}
	}

	return result
}

// Time complexity: O(n log k), since all heap operations dominate and are done up to 2k times, and k = n/3.
// Space complexity: O(n)
