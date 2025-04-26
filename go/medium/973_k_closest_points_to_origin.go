// https://leetcode.com/problems/k-closest-points-to-origin/
package medium

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
)

type PointCloset struct {
	idx  int
	dist int
}

func KClosest(points [][]int, k int) [][]int {
	minHeap := priorityqueue.NewWith(func(a, b interface{}) int {
		return b.(PointCloset).dist - a.(PointCloset).dist
	})

	for i, point := range points {
		dist := point[0]*point[0] + point[1]*point[1]
		p := PointCloset{
			idx:  i,
			dist: dist,
		}

		if minHeap.Size() > k {
			minHeap.Dequeue()
		}
		minHeap.Enqueue(p)
	}

	result := make([][]int, 0)
	for i := 0; i < k; i++ {
		p, _ := minHeap.Dequeue()
		idx := p.(PointCloset).idx
		result = append(result, points[idx])
	}

	return result
}
