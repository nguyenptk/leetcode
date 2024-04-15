// https://leetcode.com/problems/ones-and-zeroes/
package medium

import "math"

func FindMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, str := range strs {
		zeros := 0
		ones := 0
		for _, v := range str {
			if v == '0' {
				zeros++
			} else {
				ones++
			}
		}
		for i := m; i >= zeros; i-- {
			for j := n; j >= ones; j-- {
				dp[i][j] = int(math.Max(float64(dp[i][j]), float64(dp[i-zeros][j-ones]+1)))
			}
		}
	}
	return dp[m][n]
}
