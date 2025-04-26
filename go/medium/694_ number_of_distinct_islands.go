// https://leetcode.com/problems/number-of-distinct-islands/
package medium

import "fmt"

type pair struct{ r, c int }

func NumDistinctIslands(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])

	seen := make([][]bool, rows)
	for i := range seen {
		seen[i] = make([]bool, cols)
	}

	shapes := map[string]struct{}{}

	var currOrigin pair
	var currCells []pair

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return
		}
		if seen[r][c] || grid[r][c] == 0 {
			return
		}
		seen[r][c] = true
		currCells = append(currCells, pair{r - currOrigin.r, c - currOrigin.c})
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 && !seen[i][j] {
				currOrigin = pair{i, j}
				currCells = currCells[:0]
				dfs(i, j)

				key := ""
				for _, p := range currCells {
					key += fmt.Sprintf("%d_%d;", p.r, p.c)
				}
				shapes[key] = struct{}{}
			}
		}
	}

	return len(shapes)
}
