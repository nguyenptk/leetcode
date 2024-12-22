// https://leetcode.com/problems/unique-paths/
package medium

func UniquePaths(m int, n int) int {
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	return findRound(matrix, 0, 0)
}

func findRound(matrix [][]int, x, y int) int {
	if x >= len(matrix) || y >= len(matrix[0]) {
		return 0
	}
	if x == len(matrix)-1 && y == len(matrix[0])-1 {
		return 1
	}

	if matrix[x][y] > 0 {
		return matrix[x][y]
	}

	right := findRound(matrix, x+1, y)
	down := findRound(matrix, x, y+1)
	matrix[x][y] = right + down

	return matrix[x][y]
}
