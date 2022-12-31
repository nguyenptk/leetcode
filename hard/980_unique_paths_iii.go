// https://leetcode.com/problems/unique-paths-iii/
package hard

func uniquePathsIII(grid [][]int) int {
	result := 0
	empty := 1
	var sx int
	var sy int
	var ex int
	var ey int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				empty++
			} else if grid[i][j] == 1 {
				sx = i
				sy = j
			} else if grid[i][j] == 2 {
				ex = i
				ey = j
			}
		}
	}

	dfsUniquePathsIII(grid, empty, sx, sy, ex, ey, &result)

	return result
}

func dfsUniquePathsIII(grid [][]int, empty, i, j, ex, ey int, result *int) {
	if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) {
		return
	}

	if grid[i][j] < 0 {
		return
	}

	if i == ex && j == ey {
		if empty == 0 {
			*result++
		}
		return
	}

	grid[i][j] = -2
	dfsUniquePathsIII(grid, empty-1, i+1, j, ex, ey, result)
	dfsUniquePathsIII(grid, empty-1, i-1, j, ex, ey, result)
	dfsUniquePathsIII(grid, empty-1, i, j+1, ex, ey, result)
	dfsUniquePathsIII(grid, empty-1, i, j-1, ex, ey, result)
	grid[i][j] = 0
}
