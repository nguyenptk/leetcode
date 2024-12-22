// https://leetcode.com/problems/number-of-islands/
package medium

func NumIslands(grid [][]byte) int {
	result := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				result++
				dfsNumIslands(grid, i, j)
			}
		}
	}

	return result
}

func dfsNumIslands(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) {
		return
	}
	if j < 0 || j >= len(grid[0]) {
		return
	}
	if grid[i][j] == '0' {
		return
	}

	// Mark as visited so dfs does not loop
	grid[i][j] = '0'
	dfsNumIslands(grid, i-1, j)
	dfsNumIslands(grid, i+1, j)
	dfsNumIslands(grid, i, j-1)
	dfsNumIslands(grid, i, j+1)
}
