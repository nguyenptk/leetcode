// https://leetcode.com/problems/69/
package medium

var n int
var m int

func MaxAreaOfIsland(grid [][]int) int {
	result := 0
	n = len(grid)
	m = len(grid[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] > 0 {
				max := traverse(i, j, grid)
				if result < max {
					result = max
				}
			}
		}
	}

	return result
}

func traverse(i, j int, grid [][]int) int {
	if i < 0 || j < 0 || i >= n || j >= m || grid[i][j] < 1 {
		return 0
	}
	grid[i][j] = 0
	return 1 + traverse(i-1, j, grid) + traverse(i, j-1, grid) + traverse(i+1, j, grid) + traverse(i, j+1, grid)
}
