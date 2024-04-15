// https://leetcode.com/problems/task-scheduler/
package medium

import "container/heap"

type MaxHeap []byte

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(byte))
}
func (h *MaxHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

type KV struct {
	K int
	V int
}

func leastInterval(tasks []byte, n int) int {
	freq := make([]int, 26)
	for _, v := range tasks {
		freq[v-'A']++
	}

	h := &MaxHeap{}
	heap.Init(h)

	for i := 0; i < 26; i++ {
		if freq[i] > 0 {
			heap.Push(h, freq[i])
		}
	}

	time := 0
	q := make([]KV, 0)
	for h.Len() > 0 || len(q) > 0 {
		time += 1
		if h.Len() > 0 {
			val := heap.Pop(h).(int)
			val--
			if val > 0 {
				q = append(q, KV{val, time + n})
			}
		}
		if len(q) > 0 && q[0].V == time {
			heap.Push(h, q[0].K)
			q = q[1:]
		}
	}
	return time
}
