// https://leetcode.com/problems/furthest-building-you-can-reach/
package medium

import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

func FurthestBuilding(heights []int, bricks int, ladders int) int {
	h := &MinHeap{}
	heap.Init(h)
	n := len(heights) - 1
	for i := 0; i < n; i++ {
		d := heights[i+1] - heights[i]
		if d > 0 {
			heap.Push(h, d)
		}
		if h.Len() > ladders {
			bricks -= heap.Pop(h).(int)
		}
		if bricks < 0 {
			return i
		}
	}
	return n
}
