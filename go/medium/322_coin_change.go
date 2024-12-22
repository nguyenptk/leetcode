// https://leetcode.com/problems/coin-change/
package medium

func CoinChange(coins []int, amount int) int {
	// return countingTopDown(coins, amount, make(map[int]int))
	return countingBottomUp(coins, amount)
}

func countingTopDown(coins []int, amount int, memo map[int]int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	if _, ok := memo[amount]; ok {
		return memo[amount]
	}

	minCoins := -1
	for _, coin := range coins {
		subAmount := amount - coin
		subCoins := countingTopDown(coins, subAmount, memo)
		if subCoins != -1 {
			numCoins := subCoins + 1
			if minCoins == -1 || minCoins > numCoins {
				minCoins = numCoins
			}
		}
	}

	memo[amount] = minCoins

	return minCoins
}

func countingBottomUp(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 1
	for a := 1; a <= amount; a++ {
		for _, c := range coins {
			if a-c >= 0 {
				dp[a] = min(dp[a], dp[a-c]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}
