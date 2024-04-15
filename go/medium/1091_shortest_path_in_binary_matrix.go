// https://leetcode.com/problems/shortest-path-in-binary-matrix/

package medium

func ShortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] != 0 || grid[len(grid)-1][len(grid[0])-1] != 0 {
		return -1
	}

	queue := []Point{{0, 0}}
	grid[0][0] = 1

	for len(queue) > 0 {
		curr := queue[0]  // get first element
		queue = queue[1:] // remove the first element

		// if we reach the destination so we are done
		if curr.Row == len(grid)-1 && curr.Col == len(grid[0])-1 {
			return grid[curr.Row][curr.Col]
		}

		// init moves to move among points
		moves := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

		for _, move := range moves {
			t := Point{curr.Row + move[0], curr.Col + move[1]}

			if t.Row < 0 || t.Row >= len(grid) || t.Col < 0 || t.Col >= len(grid[0]) || grid[t.Row][t.Col] != 0 {
				continue
			}

			grid[t.Row][t.Col] = grid[curr.Row][curr.Col] + 1
			queue = append(queue, t)
		}
	}

	return -1
}

type Point struct {
	Row, Col int
}
