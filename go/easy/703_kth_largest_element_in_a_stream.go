// https://leetcode.com/problems/kth-largest-element-in-a-stream
package easy

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	k       int
	minHeap *IntHeap
}

func ConstructorKthLargest(k int, nums []int) KthLargest {
	tmp := IntHeap(nums)
	kthLargest := KthLargest{
		k:       k,
		minHeap: &tmp,
	}
	heap.Init(kthLargest.minHeap)
	for kthLargest.minHeap.Len() > k {
		heap.Pop(kthLargest.minHeap)
	}
	return kthLargest
}

func (t *KthLargest) Add(val int) int {
	heap.Push(t.minHeap, val)
	if t.minHeap.Len() > t.k {
		heap.Pop(t.minHeap)
	}
	return (*t.minHeap)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
