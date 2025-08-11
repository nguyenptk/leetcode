// https://leetcode.com/problems/unique-binary-search-trees/
package medium

func NumTrees(n int) int {
	if n == 0 {
		return 0
	}

	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
	}

	var rec func(start, end int) int
	rec = func(start, end int) int {
		sum := 0
		if start >= end {
			return 1
		}
		if memo[start][end] != 0 {
			return memo[start][end]
		}

		for i := start; i <= end; i++ {
			l := rec(start, i-1)
			r := rec(i+1, end)
			sum += l * r
		}
		memo[start][end] = sum
		return sum
	}

	return rec(1, n)
}

//
// func NumTrees(n int) int {
// 	if n == 0 {
// 		return 0
// 	}

// 	dp := make([]int, n+1)
// 	dp[0], dp[1] = 1, 1 // base cases

// 	for i := 2; i <= n; i++ {
// 		for j := 1; j <= i; j++ {
// 			// j is root; dp[j-1] = left subtree count, dp[i-j] = right subtree count
// 			dp[i] += dp[j-1] * dp[i-j]
// 		}
// 	}

// 	return dp[n]
// }
