// https://leetcode.com/problems/pacific-atlantic-water-flow
package medium

func PacificAtlantic(heights [][]int) [][]int {
	rows := len(heights)
	cols := len(heights[0])

	pac := make(map[int]bool)
	atl := make(map[int]bool)

	var dfs func(r, c int, visited map[int]bool, height int)
	dfs = func(r, c int, visited map[int]bool, height int) {
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return
		}
		if heights[r][c] < height {
			return
		}
		key := r*cols + c
		if visited[key] {
			return
		}
		visited[key] = true
		dfs(r+1, c, visited, heights[r][c])
		dfs(r-1, c, visited, heights[r][c])
		dfs(r, c+1, visited, heights[r][c])
		dfs(r, c-1, visited, heights[r][c])
	}

	for i := range rows {
		dfs(i, 0, pac, heights[i][0])
		dfs(i, cols-1, atl, heights[i][cols-1])
	}
	for j := range cols {
		dfs(0, j, pac, heights[0][j])
		dfs(rows-1, j, atl, heights[rows-1][j])
	}

	result := make([][]int, 0)
	for r := range rows {
		for c := range cols {
			if pac[r*cols+c] && atl[r*cols+c] {
				result = append(result, []int{r, c})
			}
		}
	}

	return result
}

// Time complexity: O(M⋅N)
// Space complexity: O(M⋅N)
