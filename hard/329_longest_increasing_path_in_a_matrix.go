// https://leetcode.com/problems/longest-increasing-path-in-a-matrix/
package hard

import "math"

// func LongestIncreasingPath(matrix [][]int) int {
// 	row := len(matrix)
// 	col := len(matrix[0])
// 	return finLongestOverAll(matrix, row, col)
// }

// // Function that returns length of the longest path
// // beginning with any cell
// func finLongestOverAll(matrix [][]int, row int, col int) int {
// 	// Initialize result
// 	result := 1

// 	// Create a lookup table and fill all entries in it as -1
// 	dp := make([][]int, row, col)
// 	for i := 0; i < row; i++ {
// 		dp[i] = make([]int, row)
// 		for j := 0; j < col; j++ {
// 			dp[i][j] = -1
// 		}
// 	}

// 	// Compute longest path beginning from all cells
// 	for i := 0; i < row; i++ {
// 		for j := 0; j < col; j++ {
// 			if dp[i][j] == -1 {
// 				findLongestFromACell(i, j, dp, matrix, row, col)
// 			}
// 			// Update result if needed
// 			result = int(math.Max(float64(result), float64(dp[i][j])))
// 		}
// 	}

// 	return result
// }

// // Return the length of LIP in 2D matrix
// func findLongestFromACell(i, j int, dp, matrix [][]int, row, col int) int {
// 	// Base case
// 	if i < 0 || i >= row || j < 0 || j >= col {
// 		return 0
// 	}

// 	// If this subproblem is already solved
// 	if dp[i][j] != -1 {
// 		return dp[i][j]
// 	}

// 	// To store the path lengths in all the four directions
// 	x := math.MinInt
// 	y := math.MinInt
// 	z := math.MinInt
// 	w := math.MinInt

// 	// Since all numbers are unique and in range from 1 to n*n,
// 	// there is atmost one possible direction from any cell
// 	if j < col-1 && ((matrix[i][j] + 1) == matrix[i][j+1]) {
// 		dp[i][j] = 1 + findLongestFromACell(i, j+1, matrix, dp, row, col)
// 		x = dp[i][j]
// 	}

// 	if j > 0 && (matrix[i][j]+1 == matrix[i][j-1]) {
// 		dp[i][j] = 1 + findLongestFromACell(i, j-1, matrix, dp, row, col)
// 		y = dp[i][j]
// 	}

// 	if i > 0 && (matrix[i][j]+1 == matrix[i-1][j]) {
// 		dp[i][j] = 1 + findLongestFromACell(i-1, j, matrix, dp, row, col)
// 		z = dp[i][j]
// 	}

// 	if i < row-1 && (matrix[i][j]+1 == matrix[i+1][j]) {
// 		dp[i][j] = 1 + findLongestFromACell(i+1, j, matrix, dp, row, col)
// 		w = dp[i][j]
// 	}

// 	// If none of the adjacent fours is one greater we will take 1
// 	// otherwise we will pick maximum from all the four directions
// 	dp[i][j] = int(math.Max(float64(x), math.Max(float64(y), math.Max(float64(z), math.Max(float64(w), 1)))))
// 	return dp[i][j]
// }

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
