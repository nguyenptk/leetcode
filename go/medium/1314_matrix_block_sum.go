// https://leetcode.com/problems/matrix-block-sum/
package medium

func MatrixBlockSum(mat [][]int, k int) [][]int {
	m := len(mat)
	n := len(mat[0])
	sum := make([][]int, m+1)
	for i := range sum {
		sum[i] = make([]int, n+1)
	}

	for r := 1; r <= m; r++ {
		for c := 1; c <= n; c++ {
			sum[r][c] = mat[r-1][c-1] + sum[r-1][c] + sum[r][c-1] - sum[r-1][c-1]
		}
	}

	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			r1, c1 := max(0, r-k), max(0, c-k)
			r2, c2 := min(m-1, r+k), min(n-1, c+k)
			r1++
			c1++
			r2++
			c2++
			result[r][c] = sum[r2][c2] - sum[r2][c1-1] - sum[r1-1][c2] + sum[r1-1][c1-1]
		}
	}

	return result
}
