// https://leetcode.com/problems/number-of-ways-to-arrive-at-destination/
package medium

import (
	"math"

	"github.com/emirpasic/gods/v2/queues/priorityqueue"
)

type Entry struct {
	dist int64 // use int64 to avoid overflow
	node int
}

func CountPaths(n int, roads [][]int) int {
	mod := 1000000007

	// Build graph: graph[node] = slice of [neighbor, weight]
	graph := make([][][2]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([][2]int, 0)
	}
	for _, road := range roads {
		start, end, time := road[0], road[1], road[2]
		graph[start] = append(graph[start], [2]int{end, time})
		graph[end] = append(graph[end], [2]int{start, time})
	}

	// Initialize shortest distances as int64 to avoid overflow.
	shortest := make([]int64, n)
	for i := 0; i < n; i++ {
		shortest[i] = math.MaxInt64
	}
	pathCount := make([]int, n)
	shortest[0] = 0
	pathCount[0] = 1

	// Priority queue with custom comparator using Entry.
	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		A := a.(Entry)
		B := b.(Entry)
		if A.dist < B.dist {
			return -1
		} else if A.dist > B.dist {
			return 1
		}
		return 0
	})
	pq.Enqueue(Entry{dist: 0, node: 0})

	for !pq.Empty() {
		top, _ := pq.Dequeue()
		entry := top.(Entry)
		currTime, currNode := entry.dist, entry.node
		if currTime > shortest[currNode] {
			continue
		}
		for _, edge := range graph[currNode] {
			neiNode, roadTime := edge[0], edge[1]
			newTime := currTime + int64(roadTime)
			if newTime < shortest[neiNode] {
				shortest[neiNode] = newTime
				pathCount[neiNode] = pathCount[currNode]
				pq.Enqueue(Entry{dist: newTime, node: neiNode})
			} else if newTime == shortest[neiNode] {
				pathCount[neiNode] = (pathCount[neiNode] + pathCount[currNode]) % mod
			}
		}
	}

	return pathCount[n-1] % mod
}
