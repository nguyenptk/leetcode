// https://leetcode.com/problems/network-delay-time/
package medium

import "math"

func NetworkDelayTime(times [][]int, n int, k int) int {
	adjList := make([][]Pair, n)
	for _, t := range times {
		adjList[t[0]-1] = append(adjList[t[0]-1], Pair{t[1], t[2]})
	}

	// Store the distance to each node in dist
	// Initialize to max
	dist := make([]int, n)
	for idx := range dist {
		dist[idx] = math.MaxInt32
	}

	queue := make([][]int, 0)
	queue = append(queue, []int{k, 0})

	visited := make(map[int]struct{})

	for len(queue) != 0 {
		item := queue[0]
		queue = queue[1:]

		curWeight := item[1]

		// Mark as visited
		visited[item[0]-1] = struct{}{}

		// No need to visited if the curweight is greater than already visited
		if dist[item[0]-1] < curWeight {
			continue
		}

		dist[item[0]-1] = curWeight

		// process adj vertices
		for _, adj := range adjList[item[0]-1] {
			// Not visited or smaller path possible
			if _, ok := visited[adj.K-1]; !ok || dist[adj.K-1] > adj.V+curWeight {
				queue = append(queue, []int{adj.K, adj.V + curWeight})
			}
		}
	}

	// If all the nodes are not visited
	if len(visited) != n {
		return -1
	}

	// Get the max time to reach a node
	max := 0
	for _, v := range dist {
		if v > max {
			max = v
		}
	}

	return max
}

type Pair struct {
	K int
	V int
}
