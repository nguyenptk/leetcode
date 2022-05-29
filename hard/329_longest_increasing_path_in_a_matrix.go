// https://leetcode.com/problems/longest-increasing-path-in-a-matrix/
package hard

import "math"

// To store the path lengths in all the four directions
var directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func LongestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	// Get row, col
	row := len(matrix)
	col := len(matrix[0])

	// Create a lookup table and fill all entries in it as 0
	dp := make([][]int, len(matrix))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}

	// Init max result
	max := 1
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			len := dfsFind(matrix, i, j, row, col, dp)
			max = int(math.Max(float64(max), float64(len)))
		}
	}
	return max
}

// Returns length of the longest path beginning with mat[i][j].
// This function mainly uses lookup table dp[row][col]
func dfsFind(matrix [][]int, i, j, row, col int, dp [][]int) int {
	// If this subproblem is already solved
	if dp[i][j] != 0 {
		return dp[i][j]
	}
	max := 1
	// Loop the directions to find the point
	for _, dir := range directions {
		x := i + dir[0]
		y := j + dir[1]
		if x < 0 || x >= row || y < 0 || y >= col || matrix[x][y] <= matrix[i][j] {
			continue
		}
		// Increse 1 unit and assign to len
		len := 1 + dfsFind(matrix, x, y, row, col, dp)
		// Find max of max & len
		max = int(math.Max(float64(max), float64(len)))
	}
	// Override the table by max
	dp[i][j] = max
	return max
}
