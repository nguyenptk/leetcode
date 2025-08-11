// https://leetcode.com/problems/rotting-oranges/
package medium

func OrangesRotting(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	fresh := 0
	q := make([][2]int, 0)

	// Initialize queue and count fresh oranges
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 2 {
				q = append(q, [2]int{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	ds := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	time := 0

	for len(q) > 0 && fresh > 0 {
		time++
		size := len(q)
		for i := 0; i < size; i++ {
			pos := q[0]
			q = q[1:]

			for _, d := range ds {
				x := pos[0] + d[0]
				y := pos[1] + d[1]

				if x >= 0 && x < rows && y >= 0 && y < cols && grid[x][y] == 1 {
					grid[x][y] = 2
					q = append(q, [2]int{x, y})
					fresh--
				}
			}
		}
	}

	if fresh == 0 {
		return time
	}
	return -1
}

// Time Complexity: O(m × n) — where m = rows, n = columns (scan + BFS)
// Space Complexity: O(m × n) in worst case (queue of all rotten oranges)
