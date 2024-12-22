// https://leetcode.com/problems/find-median-from-data-stream/
package hard

import "github.com/emirpasic/gods/queues/priorityqueue"

type MedianFinder struct {
	small *priorityqueue.Queue
	large *priorityqueue.Queue
}

func ConstructorMedianFinder() MedianFinder {
	small := priorityqueue.NewWith(func(a, b interface{}) int {
		return b.(int) - a.(int) // maxHeap
	})
	large := priorityqueue.NewWith(func(a, b interface{}) int {
		return a.(int) - b.(int) // minHeap
	})
	return MedianFinder{
		small: small,
		large: large,
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.large.Size() > 0 {
		largeTop, _ := this.large.Peek()
		if num > largeTop.(int) {
			this.large.Enqueue(num)
		} else {
			this.small.Enqueue(num)
		}
	} else {
		this.small.Enqueue(num)
	}

	// Rebalance
	if this.small.Size() > this.large.Size()+1 {
		val, _ := this.small.Dequeue()
		this.large.Enqueue(val)
	}
	if this.large.Size() > this.small.Size()+1 {
		val, _ := this.large.Dequeue()
		this.small.Enqueue(val)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.small.Size() > this.large.Size() {
		val, _ := this.small.Peek()
		return float64(val.(int))
	}
	if this.large.Size() > this.small.Size() {
		val, _ := this.large.Peek()
		return float64(val.(int))
	}
	smallVal, _ := this.small.Peek()
	largeVal, _ := this.large.Peek()
	return float64(smallVal.(int)+largeVal.(int)) / 2.0
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
