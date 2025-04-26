// https://leetcode.com/problems/minimum-path-sum/
package medium

import "math"

func MinPathSum(grid [][]int) int {
	memo := make(map[[2]int]int, 0)
	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
			return math.MaxInt
		}
		if x == len(grid)-1 && y == len(grid[0])-1 {
			return grid[x][y]
		}
		key := [2]int{x, y}
		if val, ok := memo[key]; ok {
			return val
		}

		right := dfs(x+1, y)
		down := dfs(x, y+1)

		memo[key] = grid[x][y] + min(right, down)
		return memo[key]
	}

	return dfs(0, 0)
}
