// https://leetcode.com/problems/out-of-boundary-paths/
package medium

func FindPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	if maxMove == 0 {
		return 0
	}

	dpCurr := make([][]int, m+2)
	for i := range dpCurr {
		dpCurr[i] = make([]int, n+2)
	}

	dpLast := make([][]int, m+2)
	for i := range dpLast {
		dpLast[i] = make([]int, n+2)
	}

	for i := 1; i <= m; i++ {
		dpCurr[i][1]++
		dpCurr[i][n]++
	}
	for j := 1; j <= n; j++ {
		dpCurr[1][j]++
		dpCurr[m][j]++
	}

	result := dpCurr[startRow+1][startColumn+1]
	for d := 1; d < maxMove; d++ {
		temp := dpCurr
		dpCurr = dpLast
		dpLast = temp
		for i := 1; i <= m; i++ {
			for j := 1; j <= n; j++ {
				dpCurr[i][j] = (dpLast[i-1][j] + dpLast[i+1][j] + dpLast[i][j-1] + dpLast[i][j+1]) % 1000000007
			}
		}

		result = (result + dpCurr[startRow+1][startColumn+1]) % 1000000007
	}
	return result
}
