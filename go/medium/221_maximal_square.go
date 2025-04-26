// https://leetcode.com/problems/maximal-square/
package medium

func maximalSquare(matrix [][]byte) int {
	rows := len(matrix)
	cols := len(matrix[0])

	dp := make([][]int, rows+1)
	for i := range dp {
		dp[i] = make([]int, cols+1)
	}

	maxSide := 0
	for i := 1; i < rows+1; i++ {
		for j := 1; j < cols+1; j++ {
			if matrix[i-1][j-1] == '1' {
				top := int(dp[i-1][j])
				left := int(dp[i][j-1])
				diag := int(dp[i-1][j-1])
				dp[i][j] = min(top, min(left, diag)) + 1
				maxSide = max(maxSide, dp[i][j])
			}
		}
	}

	return maxSide * maxSide
}
