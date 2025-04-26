// https://leetcode.com/problems/most-profitable-path-in-a-tree/
package medium

import "math"

func MostProfitablePath(edges [][]int, bob int, amount []int) int {
	n := len(amount)
	adj := make(map[int][]int)

	// Build the adjacency list
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	// Initialize bobDistance with a large number
	bobDistance := make([]int, n)
	for i := range bobDistance {
		bobDistance[i] = n
	}

	// Internal DFS function
	var dfs func(src, parent, time int) int
	dfs = func(src, parent, time int) int {
		maxIncome := 0
		maxChild := math.MinInt64 // Use MinInt64 for initialization

		// Set distance from Bob
		if src == bob {
			bobDistance[src] = 0
		}

		// Traverse the adjacent nodes
		for _, nei := range adj[src] {
			if nei != parent {
				childProfit := dfs(nei, src, time+1)
				maxChild = max(maxChild, childProfit)
				bobDistance[src] = min(bobDistance[src], bobDistance[nei]+1)
			}
		}

		// Calculate income based on Alice's and Bob's arrival time
		if bobDistance[src] > time {
			maxIncome += amount[src]
		} else if bobDistance[src] == time {
			maxIncome += amount[src] / 2
		}

		// If it's a leaf node
		if maxChild == math.MinInt64 {
			return maxIncome
		}

		return maxIncome + maxChild
	}

	return dfs(0, -1, 0)
}
