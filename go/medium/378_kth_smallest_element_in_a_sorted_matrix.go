// https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/
package medium

import "container/heap"

type T struct {
	matrix *[][]int
	row    int
	col    int
}

type minHeap []T

func KthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}

	h := &minHeap{}

	heap.Push(h, T{&matrix, 0, 0})
	for i := 1; i < n; i++ {
		heap.Push(h, T{&matrix, i, 0})
		heap.Push(h, T{&matrix, 0, i})
	}

	var res T

	set := make(map[T]struct{})

	for ; k > 0; k-- {
		res = heap.Pop(h).(T)

		var newT T
		if res.row == res.col {
			continue
		} else if res.row > res.col {
			newT = T{&matrix, res.row, res.col + 1}
		} else {
			newT = T{&matrix, res.row + 1, res.col}
		}

		if _, ok := set[newT]; !ok {
			set[newT] = struct{}{}
			heap.Push(h, newT)
		}
	}
	return matrix[res.row][res.col]
}

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i int, j int) bool {
	return (*(h[i].matrix))[h[i].row][h[i].col] < (*(h[j].matrix))[h[j].row][h[j].col]
}

func (h minHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(a interface{}) {
	*h = append(*h, a.(T))
}

func (h *minHeap) Pop() interface{} {
	l := len(*h)
	res := (*h)[l-1]
	*h = (*h)[:l-1]
	return res
}
