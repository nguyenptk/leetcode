// https://leetcode.com/problems/delete-operation-for-two-strings/
package medium

import "math"

func MinDistance(word1 string, word2 string) int {
	if word1 == word2 {
		return 0
	}

	l1 := len(word1)
	l2 := len(word2)

	return findDiff(word1, word2, l1, l2)
}

func findDiff(word1, word2 string, l1, l2 int) int {
	// init dp
	dp := make([][]int, l1+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, l2+1)
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = maxNums(1+dp[i-1][j-1], dp[i-1][j], dp[i][j-1])
			} else {
				dp[i][j] = maxNums(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[l1][l2]
}

func maxNums(nums ...int) int {
	res := math.MinInt64
	for _, num := range nums {
		if num > res {
			res = num
		}
	}
	return res
}
