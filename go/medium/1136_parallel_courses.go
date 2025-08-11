// https://leetcode.com/problems/parallel-courses/
package medium

func MinimumSemesters(n int, relations [][]int) int {
	graph := make(map[int][]int)
	for _, r := range relations {
		u, v := r[0], r[1]
		graph[u] = append(graph[u], v)
	}

	visited := make([]int, n+1)
	var dfs func(idx int) int
	dfs = func(idx int) int {
		if visited[idx] != 0 {
			return visited[idx]
		}
		visited[idx] = -1
		maxLength := 1
		for _, end := range graph[idx] {
			length := dfs(end)
			if length < 0 {
				return -1
			}
			maxLength = max(maxLength, length+1)
		}
		visited[idx] = maxLength

		return maxLength
	}

	maxLength := 1
	for i := 1; i <= n; i++ {
		length := dfs(i)
		if length < 0 {
			return -1
		}
		maxLength = max(maxLength, length)
	}

	return maxLength
}
