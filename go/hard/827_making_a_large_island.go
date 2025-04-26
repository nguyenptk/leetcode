// https://leetcode.com/problems/making-a-large-island/
package hard

func LargestIsland(grid [][]int) int {
	islandSizes := map[int]int{}
	islandID := 2

	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] != 1 {
			return 0
		}
		grid[row][col] = islandID
		return 1 + dfs(row+1, col) + dfs(row-1, col) + dfs(row, col+1) + dfs(row, col-1)
	}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == 1 {
				islandSizes[islandID] = dfs(row, col)
				islandID++
			}
		}
	}

	// Edge case: If all cells are 1
	allOnes := true
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == 0 {
				allOnes = false
				break
			}
		}
		if !allOnes {
			break
		}
	}
	if allOnes {
		return len(grid) * len(grid[0])
	}

	maxIslandSize := 1
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == 0 {
				curIslandSize := 1
				neiIslands := map[int]bool{}

				if row+1 < len(grid) && grid[row+1][col] > 1 {
					neiIslands[grid[row+1][col]] = true
				}
				if row-1 >= 0 && grid[row-1][col] > 1 {
					neiIslands[grid[row-1][col]] = true
				}
				if col+1 < len(grid[0]) && grid[row][col+1] > 1 {
					neiIslands[grid[row][col+1]] = true
				}
				if col-1 >= 0 && grid[row][col-1] > 1 { // Fixing column check
					neiIslands[grid[row][col-1]] = true
				}

				for id := range neiIslands {
					curIslandSize += islandSizes[id]
				}
				maxIslandSize = max(maxIslandSize, curIslandSize)
			}
		}
	}

	return maxIslandSize
}
