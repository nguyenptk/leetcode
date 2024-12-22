// https://leetcode.com/problems/longest-common-subsequence/
package medium

func LongestCommonSubsequence(text1 string, text2 string) int {
	// return topDownLongestCommonSubsequence(text1, text2)
	return bottomUpLongestCommonSubsequence(text1, text2)
}

func topDownLongestCommonSubsequence(text1 string, text2 string) int {
	memo := make([][]int, len(text1))
	for i := range memo {
		memo[i] = make([]int, len(text2))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return solveTopDown(text1, text2, 0, 0, memo)
}

func solveTopDown(text1, text2 string, x, y int, memo [][]int) int {
	count := 0
	if x >= len(text1) {
		return 0
	}
	if y >= len(text2) {
		return 0
	}

	if memo[x][y] != -1 {
		return memo[x][y]
	}

	if text1[x] == text2[y] {
		count = 1 + solveTopDown(text1, text2, x+1, y+1, memo)
	} else {
		count = max(solveTopDown(text1, text2, x+1, y, memo),
			solveTopDown(text1, text2, x, y+1, memo),
		)
	}

	memo[x][y] = count
	return count
}

func bottomUpLongestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := len(text1) - 1; i >= 0; i-- {
		for j := len(text2) - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return dp[0][0]
}
