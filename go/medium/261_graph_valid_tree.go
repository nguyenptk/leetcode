// https://leetcode.com/problems/graph-valid-tree
package medium

func ValidTreeDFS(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}

	graph := buildGraph(edges)
	seen := make(map[int]bool)

	var dfs func(node int)
	dfs = func(node int) {
		if seen[node] {
			return
		}
		seen[node] = true

		for _, neighbor := range graph[node] {
			dfs(neighbor)
		}
	}

	dfs(0)

	return len(seen) == n
}

func ValidTreeBFS(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}

	graph := buildGraph(edges)
	seen := make(map[int]bool)
	queue := make([]int, 0)
	queue = append(queue, 0)
	seen[0] = true

	for len(queue) > 0 {
		node := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		for _, neighbor := range graph[node] {
			if seen[neighbor] {
				continue
			}
			seen[neighbor] = true
			queue = append(queue, neighbor)
		}
	}

	return len(seen) == n
}

func buildGraph(edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	for _, e := range edges {
		x := e[0]
		y := e[1]
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}

	return graph
}
