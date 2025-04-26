// https://leetcode.com/problems/01-matrix/
package medium

func UpdateMatrix(mat [][]int) [][]int {
	rows := len(mat)
	cols := len(mat[0])

	result := make([][]int, rows)
	for i := range rows {
		result[i] = make([]int, cols)
		for j := range cols {
			result[i][j] = -1
		}
	}

	q := make([][]int, 0)
	for i := range rows {
		for j := range cols {
			if mat[i][j] == 0 {
				q = append(q, []int{i, j})
				result[i][j] = 0
			}
		}
	}

	ds := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	dist := 1
	for len(q) > 0 {
		length := len(q)
		for i := 0; i < length; i++ {
			ceil := q[0]
			q = q[1:]
			for _, d := range ds {
				x := d[0] + ceil[0]
				y := d[1] + ceil[1]
				if x >= 0 && x < rows && y >= 0 && y < cols && result[x][y] == -1 {
					result[x][y] = dist
					q = append(q, []int{x, y})
				}
			}
		}
		dist++
	}

	return result
}
