package medium

func CountComponents(n int, edges [][]int) int {
	result := 0

	graph := createGraph(edges)
	visited := make(map[int]bool, n)

	var dfs func(int)
	dfs = func(node int) {
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				dfs(neighbor)
			}
		}
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			visited[i] = true
			dfs(i)
			result++
		}
	}

	return result
}

func createGraph(edges [][]int) map[int][]int {
	graph := make(map[int][]int, 0)
	for _, e := range edges {
		x := e[0]
		y := e[1]
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}

	return graph
}
