// https://leetcode.com/problems/count-the-number-of-complete-components/
package medium

func CountCompleteComponents(n int, edges [][]int) int {
	graph := make([][]int, n)

	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	visited := map[int]bool{}
	var dfs func(idx int, compInfo []int)
	dfs = func(idx int, compInfo []int) {
		visited[idx] = true
		compInfo[0]++
		compInfo[1] += len(graph[idx])

		for _, nei := range graph[idx] {
			if !visited[nei] {
				dfs(nei, compInfo)
			}
		}
	}

	result := 0
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		compInfo := make([]int, 2)
		dfs(i, compInfo)

		if compInfo[0]*(compInfo[0]-1) == compInfo[1] {
			result++
		}
	}

	return result
}
