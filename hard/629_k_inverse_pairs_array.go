// https://leetcode.com/problems/k-inverse-pairs-array/
package hard

func KInversePairs(n int, k int) int {
	// Initialize 2D DP
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	// There if no inverse pair, the permutation is unique "123..i"
	for i := 0; i <= n; i++ {
		dp[i][0] = 1
	}

	// Else, find the inverse pair
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			dp[i][j] = (dp[i][j-1] + dp[i-1][j]) % (1e9 + 7)
			if j-i >= 0 {
				dp[i][j] = (dp[i][j] - dp[i-1][j-i] + (1e9 + 7)) % (1e9 + 7)
			}

		}
	}

	return dp[n][k]
}
