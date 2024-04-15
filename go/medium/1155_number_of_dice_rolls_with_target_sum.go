// https://leetcode.com/problems/number-of-dice-rolls-with-target-sum/
package medium

func NumRollsToTarget(n int, k int, target int) int {
	kMod := 1000000007
	dp := make([]int, target+1)
	dp[0] = 1

	for ; n > 0; n-- {
		newDP := make([]int, target+1)
		for i := 1; i <= k; i++ {
			for t := i; t <= target; t++ {
				newDP[t] += dp[t-i]
				newDP[t] %= kMod
			}
		}
		dp = newDP
	}

	return dp[target]
}
