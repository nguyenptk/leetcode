// https://leetcode.com/problems/maximum-performance-of-a-team/
package hard

import (
	"container/heap"
	"sort"
)

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

func MaxPerformance(n int, speed []int, efficiency []int, k int) int {
	kMod := (int)(1e9 + 7)

	ord := make([][]int, n)
	for i := range ord {
		ord[i] = []int{efficiency[i], speed[i]}
	}

	sort.Slice(ord, func(i, j int) bool {
		return ord[i][0] > ord[j][0]
	})

	h := &MinHeap{}
	heap.Init(h)

	totalSpeed := 0
	best := 0

	for _, v := range ord {
		if h.Len() >= k {
			totalSpeed -= heap.Pop(h).(int)
		}
		totalSpeed += v[1]
		heap.Push(h, v[1])

		perf := totalSpeed * v[0]
		if best < perf {
			best = perf
		}
	}

	return best % kMod
}
