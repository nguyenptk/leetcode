// https://leetcode.com/problems/minimum-difficulty-of-a-job-schedule/
package hard

import "math"

func MinDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if n < d {
		return -1
	}

	// dp[i][k]: min difficulty to schedule the first i jobs in k days
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, d+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt / 2
		}
	}
	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		for k := 1; k <= d; k++ {
			maxDifficulty := 0              // Max(job[j + 1..i])
			for j := i - 1; j >= k-1; j-- { // 1-based
				if maxDifficulty < jobDifficulty[j] { // 0-based
					maxDifficulty = jobDifficulty[j]
				}
				if dp[i][k] > dp[j][k-1]+maxDifficulty {
					dp[i][k] = dp[j][k-1] + maxDifficulty
				}
			}
		}
	}

	return dp[n][d]
}
