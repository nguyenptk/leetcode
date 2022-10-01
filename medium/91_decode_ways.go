// https://leetcode.com/problems/decode-ways/
package medium

func NumDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}

	n := len(s)
	dp := make([]int, n+1)
	// base cases
	dp[n] = 1
	if isValidOne(s[n-1]) {
		dp[n-1] = 1
	}

	for i := n - 2; i >= 0; i-- {
		if isValidOne(s[i]) {
			dp[i] += dp[i+1]
		}
		if isValidTwo(s[i], s[i+1]) {
			dp[i] += dp[i+2]
		}
	}

	return dp[0]
}

func isValidOne(c byte) bool {
	return c != '0'
}

func isValidTwo(c1, c2 byte) bool {
	return c1 == '1' || c1 == '2' && c2 < '7'
}
