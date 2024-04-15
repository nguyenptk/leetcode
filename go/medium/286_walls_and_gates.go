// https://leetcode.com/problems/walls-and-gates
package medium

import (
	"math"
)

func WallsAndGates(rooms [][]int) {
	const (
		EMPTY = math.MaxInt32
		GATE  = 0
	)

	m := len(rooms)
	if m == 0 {
		return
	}
	n := len(rooms[0])
	q := make([][]int, 0)
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if rooms[row][col] == GATE {
				q = append(q, []int{row, col})
			}
		}
	}
	directions := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	for len(q) > 0 {
		cell := q[0]
		q = q[1:]
		row := cell[0]
		col := cell[1]
		for _, d := range directions {
			r := row + d[0]
			c := col + d[1]
			if r < 0 || c < 0 || r >= m || c >= n || rooms[r][c] != EMPTY {
				continue
			}
			rooms[r][c] = rooms[row][col] + 1
			q = append(q, []int{r, c})
		}
	}
}
