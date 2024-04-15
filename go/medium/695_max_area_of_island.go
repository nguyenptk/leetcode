// https://leetcode.com/problems/69/
package medium

func MaxAreaOfIsland(grid [][]int) int {
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				result = max(result, dfsMaxAreaOfIsland(grid, i, j))
			}
		}
	}
	return result
}

func dfsMaxAreaOfIsland(grid [][]int, i, j int) int {
	if i < 0 || i >= len(grid) {
		return 0
	}
	if j < 0 || j >= len(grid[0]) {
		return 0
	}
	if grid[i][j] == 0 {
		return 0
	}

	// Marked as visited
	grid[i][j] = 0
	return 1 + dfsMaxAreaOfIsland(grid, i+1, j) + dfsMaxAreaOfIsland(grid, i-1, j) + dfsMaxAreaOfIsland(grid, i, j+1) + dfsMaxAreaOfIsland(grid, i, j-1)
}
