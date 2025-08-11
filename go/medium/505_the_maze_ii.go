// https://leetcode.com/problems/the-maze-ii/
package medium

import "math"

func ShortestDistance(maze [][]int, start []int, destination []int) int {
	m, n := len(maze), len(maze[0])
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt
		}
	}

	sx, sy := start[0], start[1]
	dx, dy := destination[0], destination[1]
	dist[sx][sy] = 0

	var dfs func(x, y int)
	dfs = func(x, y int) {
		dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		for _, d := range dirs {
			nx, ny, steps := x+d[0], y+d[1], 0
			for nx >= 0 && nx < m &&
				ny >= 0 && ny < n &&
				maze[nx][ny] == 0 {
				nx += d[0]
				ny += d[1]
				steps++
			}
			stopX := nx - d[0]
			stopY := ny - d[1]

			if dist[x][y]+steps < dist[stopX][stopY] {
				dist[stopX][stopY] = dist[x][y] + steps
				dfs(stopX, stopY)
			}
		}
	}

	dfs(sx, sy)
	if dist[dx][dy] == math.MaxInt {
		return -1
	}
	return dist[dx][dy]
}
