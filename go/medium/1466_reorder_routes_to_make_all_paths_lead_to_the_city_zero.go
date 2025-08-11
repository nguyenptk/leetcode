// https://leetcode.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/
package medium

func MinReorder(n int, connections [][]int) int {
	result := 0
	graph := make(map[int][][]int, 0)
	for _, c := range connections {
		u, v := c[0], c[1]
		graph[u] = append(graph[u], []int{v, 1})
		graph[v] = append(graph[v], []int{u, 0})
	}

	var dfs func(node int, parent int)
	dfs = func(node int, parent int) {
		for _, nei := range graph[node] {
			neighbor := nei[0]
			sign := nei[1]
			if neighbor != parent {
				result += sign
				dfs(neighbor, node)
			}
		}
	}

	dfs(0, -1)
	return result
}
