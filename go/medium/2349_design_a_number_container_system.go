// https://leetcode.com/problems/design-a-number-container-system/
package medium

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
)

type NumberContainers struct {
	idxNums     map[int]int
	numIndicies map[int]*priorityqueue.Queue
}

func ConstructorNumberContainers() NumberContainers {
	return NumberContainers{
		idxNums:     map[int]int{},
		numIndicies: map[int]*priorityqueue.Queue{},
	}
}

func (f *NumberContainers) Change(index int, number int) {
	// Remove from old number queue
	oldNum, ok := f.idxNums[index]
	if ok && oldNum != number {
		if queue, ok := f.numIndicies[oldNum]; ok {
			newQueue := priorityqueue.NewWith(func(a, b interface{}) int {
				return a.(int) - b.(int)
			})
			for !queue.Empty() {
				item, _ := queue.Dequeue()
				if item.(int) != index {
					newQueue.Enqueue(item)
				}
			}
			// Remove empty queue
			if newQueue.Empty() {
				delete(f.numIndicies, oldNum)
			} else {
				f.numIndicies[oldNum] = newQueue
			}
		}
	}

	f.idxNums[index] = number
	if _, ok := f.numIndicies[number]; !ok {
		f.numIndicies[number] = priorityqueue.NewWith(func(a, b interface{}) int {
			return a.(int) - b.(int)
		})
	}

	f.numIndicies[number].Enqueue(index)
}

func (f *NumberContainers) Find(number int) int {
	if queue, ok := f.numIndicies[number]; ok {
		for !queue.Empty() {
			smallest, _ := queue.Peek()
			if val, exists := f.idxNums[smallest.(int)]; exists && val == number {
				return smallest.(int)
			}
			queue.Dequeue() // Remove stale index
		}
	}
	return -1
}

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */
