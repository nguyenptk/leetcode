// https://leetcode.com/problems/optimal-account-balancing/
package hard

import (
	"fmt"
	"math"
	"strings"
)

func MinTransfers(transactions [][]int) int {
	balanceM := make(map[int]int)
	for _, t := range transactions {
		from, to, amount := t[0], t[1], t[2]
		balanceM[from] -= amount
		balanceM[to] += amount
	}

	balanceL := make([]int, 0)
	for _, amt := range balanceM {
		if amt != 0 {
			balanceL = append(balanceL, amt)
		}
	}

	n := len(balanceL)
	memo := make(map[string]int)

	var dfs func(int) int
	dfs = func(start int) int {
		for start < n && balanceL[start] == 0 {
			start++
		}
		if start == n {
			return 0
		}

		key := hash(balanceL, start)
		if val, ok := memo[key]; ok {
			return val
		}

		cost := math.MaxInt
		for i := start + 1; i < n; i++ {
			if balanceL[i]*balanceL[start] < 0 {
				// try to settle balance[start] with balance[i]
				balanceL[i] += balanceL[start]
				cost = min(cost, 1+dfs(start+1))
				balanceL[i] -= balanceL[start]
			}
		}
		memo[key] = cost
		return memo[key]
	}

	return dfs(0)
}

func hash(b []int, start int) string {
	var sb strings.Builder
	for i := start; i < len(b); i++ {
		sb.WriteString(fmt.Sprintf("%d,", b[i]))
	}
	return sb.String()
}

// Time: Worst-case O(n!), but in practice much better due to memoization and pruning.
// Space: O(n) for balanceL
// O(n!) max memo states in worst case
// Each memo key is O(n) in size

// -----------------------------

// Optimal solution with bitmask
// func minTransfers(transactions [][]int) int {
// 	balanceMap := make(map[int]int)
// 	for _, t := range transactions {
// 		from, to, amount := t[0], t[1], t[2]
// 		balanceMap[from] += amount
// 		balanceMap[to] -= amount
// 	}

// 	credits := []int{}
// 	for _, v := range balanceMap {
// 		if v != 0 {
// 			credits = append(credits, v)
// 		}
// 	}

// 	n := len(credits)
// 	memo := make([]int, 1<<n) // 1<<n = 2^n
// 	for i := range memo {
// 		memo[i] = -1
// 	}
// 	memo[0] = 0

// 	var dfs func(mask int) int
// 	dfs = func(mask int) int {
// 		if memo[mask] != -1 {
// 			return memo[mask]
// 		}
// 		sum, best := 0, 0
// 		for i := 0; i < n; i++ {
// 			if mask&(1<<i) != 0 {
// 				sum += credits[i]
// 				best = max(best, dfs(mask^(1<<i)))
// 			}
// 		}
// 		if sum == 0 {
// 			memo[mask] = best + 1
// 		} else {
// 			memo[mask] = best
// 		}
// 		return memo[mask]
// 	}

// 	return n - dfs((1<<n)-1)
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// // Time complexity: O(n⋅2^n)
// // We build memo, an array of size O(2^n) as memory, equal to the number of possible states. Each state is computed with a traverse through balance_list, which takes O(n) time.
// // Space complexity: O(2^n)
// // The length of memo is 2^n
// // The space complexity of a recursive call depends on the maximum depth of the recursive call stack, which is n. As each recursive call removes one set bit from total_mask. Therefore, at most O(n) levels of recursion will be created, and each level consumes a constant amount of space.
