// https://leetcode.com/problems/rotate-image/
package medium

// Add additional return value for unit-test
func Rotate(matrix [][]int) [][]int {
	n := len(matrix)
	for i := 0; i < (n+1)/2; i++ {
		for j := 0; j < (n / 2); j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[n-j-1][i]
			matrix[n-j-1][i] = matrix[n-i-1][n-j-1]
			matrix[n-i-1][n-j-1] = matrix[j][n-i-1]
			matrix[j][n-i-1] = tmp
		}
	}
	return matrix
}
