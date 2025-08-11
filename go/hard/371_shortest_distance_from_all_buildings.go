// https://leetcode.com/problems/shortest-distance-from-all-buildings/
package hard

import "math"

func ShortestDistance(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	total := make([][]int, rows)
	for i := range total {
		total[i] = make([]int, cols)
	}

	emptyLandValue := 0
	minDist := math.MaxInt
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				minDist = math.MaxInt
				q := make([][2]int, 0)
				q = append(q, [2]int{r, c})
				steps := 0

				for len(q) > 0 {
					steps++
					for level := len(q); level > 0; level-- {
						curr := q[0]
						q = q[1:]

						for _, d := range dirs {
							nextR := curr[0] + d[0]
							nextC := curr[1] + d[1]

							if nextR >= 0 && nextR < rows &&
								nextC >= 0 && nextC < cols &&
								grid[nextR][nextC] == emptyLandValue {
								grid[nextR][nextC]--
								total[nextR][nextC] += steps
								q = append(q, [2]int{nextR, nextC})
								minDist = min(minDist, total[nextR][nextC])
							}
						}
					}
				}
				emptyLandValue--
			}
		}
	}

	if minDist == math.MaxInt {
		return -1
	}
	return minDist
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
