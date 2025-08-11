// https://leetcode.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/
package hard

type StatePath struct {
	row   int
	col   int
	k     int
	steps int
}

func ShortestPath(grid [][]int, k int) int {
	rows, cols := len(grid), len(grid[0])
	if k >= rows+cols-2 {
		return rows + cols - 2
	}

	start := StatePath{0, 0, k, 0}
	queue := []StatePath{start}
	visited := map[[3]int]bool{
		{0, 0, k}: true,
	}

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.row == rows-1 && curr.col == cols-1 {
			return curr.steps
		}

		for _, d := range dirs {
			nr, nc := curr.row+d[0], curr.col+d[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
				nk := curr.k - grid[nr][nc]
				key := [3]int{nr, nc, nk}
				if nk >= 0 && !visited[key] {
					visited[key] = true
					queue = append(queue, StatePath{nr, nc, nk, curr.steps + 1})
				}
			}
		}
	}
	return -1
}
