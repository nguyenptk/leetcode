package medium

import "math"

func MinCostConnectPoints(points [][]int) int {
	n := len(points)
	node := 0
	dist := make([]int, n)
	for i := range dist {
		dist[i] = 100000000
	}
	visit := make([]bool, n)
	edges, result := 0, 0

	for edges < n-1 {
		visit[node] = true
		nextNode := -1
		for i := 0; i < n; i++ {
			if visit[i] {
				continue
			}
			curDist := int(math.Abs(float64(points[i][0]-points[node][0])) + math.Abs(float64(points[i][1]-points[node][1])))
			if curDist < dist[i] {
				dist[i] = curDist
			}
			if nextNode == -1 || dist[i] < dist[nextNode] {
				nextNode = i
			}
		}
		result += dist[nextNode]
		node = nextNode
		edges++
	}

	return result
}
