// https://leetcode.com/problems/number-of-distinct-islands/
package medium

import "strings"

func NumDistinctIslands(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	unique := make(map[string]bool)
	visited := make(map[int]bool)

	var path strings.Builder
	var dfs func(row, col int, dir string)
	dfs = func(row, col int, dir string) {
		if row < 0 || row >= rows ||
			col < 0 || col >= cols {
			return
		}

		key := row*cols + col
		if visited[key] || grid[row][col] != 1 {
			return
		}
		visited[key] = true

		path.WriteString(dir)

		dfs(row+1, col, "D")
		dfs(row-1, col, "U")
		dfs(row, col+1, "R")
		dfs(row, col-1, "L")

		path.WriteString("0")
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			path.Reset()
			dfs(row, col, "0")
			if path.Len() > 0 {
				unique[path.String()] = true
			}
		}
	}

	return len(unique)
}

// Time Complexity: O(M⋅N). Same as Approach 2.
// Space complexity: O(M⋅N). Same as Approach 2.
