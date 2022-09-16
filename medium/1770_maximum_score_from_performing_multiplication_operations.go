// https://leetcode.com/problems/maximum-score-from-performing-multiplication-operations/
package medium

var dp [][]int

func MaximumScore(nums []int, multipliers []int) int {
	n := len(multipliers)
	// dp[s][i] := max score of nums[s..e] and multipliers[i],
	dp = make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	return maxScore(nums, 0, multipliers, 0, n)
}

func maxScore(nums []int, s int, multipliers []int, i int, n int) int {
	if i == n {
		return 0
	}

	if dp[s][i] != 0 {
		return dp[s][i]
	}

	// # of nums picked on the start side = s,
	// # of nums picked on the end side = i - s,
	// so e = n - (i - s) - 1
	e := len(nums) - (i - s) - 1
	pickStart := nums[s]*multipliers[i] + maxScore(nums, s+1, multipliers, i+1, n)
	pickEnd := nums[e]*multipliers[i] + maxScore(nums, s, multipliers, i+1, n)

	if pickStart < pickEnd {
		dp[s][i] = pickEnd
	} else {
		dp[s][i] = pickStart
	}

	return dp[s][i]
}
