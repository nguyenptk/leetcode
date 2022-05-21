// https://leetcode.com/problems/coin-change/
package medium

import "math"

func CoinChange(coins []int, amount int) int {

	// table[i] will be storing the number of solutions for
	// value i. We need n+1 rows as the table is constructed
	// in bottom up manner using the base case (n = 0)
	table := make([]int, amount+1)

	// Base case (If given value is 0)
	table[0] = 0

	// Initialize all table values as Infinite
	for i := 1; i <= amount; i++ {
		table[i] = math.MaxInt
	}

	// Compute minimum coins required for all
	// values from 1 to amount
	for i := 1; i <= amount; i++ {
		// Go through all coins smaller than i
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				subRes := table[i-coins[j]]
				if subRes != math.MaxInt && subRes+1 < table[i] {
					table[i] = subRes + 1
				}
			}
		}
	}

	if table[amount] == math.MaxInt {
		return -1
	}

	return table[amount]
}
