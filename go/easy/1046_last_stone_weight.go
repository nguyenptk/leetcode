// https://leetcode.com/problems/last-stone-weight/
package easy

import "github.com/emirpasic/gods/queues/priorityqueue"

func LastStoneWeight(stones []int) int {
	heap := priorityqueue.NewWith(func(a, b interface{}) int {
		return b.(int) - a.(int)
	})

	for _, stone := range stones {
		heap.Enqueue(stone)
	}

	for heap.Size() > 1 {
		y, _ := heap.Dequeue()
		x, _ := heap.Dequeue()
		if x != y {
			newVal := y.(int) - x.(int)
			heap.Enqueue(newVal)
		}
	}

	result, _ := heap.Dequeue()
	if result == nil {
		return 0
	}

	return result.(int)
}
