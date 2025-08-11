// https://leetcode.com/problems/69/
package medium

func MaxAreaOfIsland(grid [][]int) int {
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= len(grid) ||
			j < 0 || j >= len(grid[0]) ||
			grid[i][j] == 0 {
			return 0
		}
		grid[i][j] = 0
		return 1 + dfs(i+1, j) + dfs(i-1, j) + dfs(i, j+1) + dfs(i, j-1)
	}

	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				result = max(result, dfs(i, j))
			}
		}
	}
	return result
}

// Time Complexity: O(nm)
// Space complexity: O(nm)
