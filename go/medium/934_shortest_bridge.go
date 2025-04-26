// https://leetcode.com/problems/shortest-bridge/
package medium

func ShortestBridge(grid [][]int) int {
	n := len(grid)
	firstX, firstY := -1, -1
	for i := range n {
		for j := range n {
			if grid[i][j] == 1 {
				firstX = i
				firstY = j
				break
			}
		}
	}

	queue := make([][2]int, 0)
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= n || y < 0 || y >= n || grid[x][y] != 1 {
			return
		}
		grid[x][y] = 2
		queue = append(queue, [2]int{x, y})
		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y+1)
		dfs(x, y-1)
	}
	dfs(firstX, firstY)

	distance := 0
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(queue) > 0 {
		newQueue := make([][2]int, 0)
		for _, pair := range queue {
			x := pair[0]
			y := pair[1]
			for _, d := range dirs {
				curX := x + d[0]
				curY := y + d[1]
				if curX < 0 || curX >= n || curY < 0 || curY >= n {
					continue
				}
				if grid[curX][curY] == 1 {
					return distance
				}
				if grid[curX][curY] == 0 {
					grid[curX][curY] = 2
					newQueue = append(newQueue, [2]int{curX, curY})
				}
			}
		}
		queue = newQueue
		distance++
	}

	return -1
}
