package medium

type NumMatrix struct {
	dp [][]int
}

func Constructor2D(matrix [][]int) NumMatrix {
	yLen := len(matrix) + 1
	xLen := len(matrix[0]) + 1
	dp := make([][]int, yLen)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, xLen)
	}
	for i := 1; i < yLen; i++ {
		for j := 1; j < xLen; j++ {
			dp[i][j] = matrix[i-1][j-1] + dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1]
		}
	}

	return NumMatrix{
		dp: dp,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.dp[row2+1][col2+1] - this.dp[row2+1][col1] - this.dp[row1][col2+1] + this.dp[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
