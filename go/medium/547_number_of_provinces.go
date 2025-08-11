// https://leetcode.com/problems/number-of-provinces/
package medium

func FindCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	result := 0
	visited := make([]bool, n)

	var dfs func(node int)
	dfs = func(node int) {
		visited[node] = true
		for i := 0; i < n; i++ {
			if isConnected[node][i] == 1 && !visited[i] {
				dfs(i)
			}
		}
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			result++
			dfs(i)
		}
	}

	return result
}

// Time complexity: O(n^2)
// Space complexity: O(n)
