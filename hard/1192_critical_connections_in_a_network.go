// https://leetcode.com/problems/critical-connections-in-a-network/
package hard

import "math"

var disc, low []int
var time int = 0

var ans [][]int
var edgeMap map[int][]int = map[int][]int{}

func CriticalConnections(n int, connections [][]int) [][]int {
	disc = make([]int, n)
	low = make([]int, n)
	for i := 0; i < n; i++ {
		edgeMap[i] = []int{}
	}
	for _, conn := range connections {
		edgeMap[conn[0]] = append(edgeMap[conn[0]], conn[1])
		edgeMap[conn[1]] = append(edgeMap[conn[1]], conn[0])
	}
	ans = [][]int{}
	dfs(0, -1)
	return ans
}

func dfs(curr, prev int) {
	time++
	disc[curr] = time
	low[curr] = time

	for _, next := range edgeMap[curr] {
		if disc[next] == 0 {
			dfs(next, curr)
			low[curr] = int(math.Min(float64(low[curr]), float64(low[next])))
		} else if next != prev {
			low[curr] = int(math.Min(float64(low[curr]), float64(disc[next])))
		}
		if low[next] > disc[curr] {
			ans = append(ans, [][]int{{curr, next}}...)
		}
	}
}
