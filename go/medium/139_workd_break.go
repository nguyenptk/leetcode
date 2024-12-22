// https://leetcode.com/problems/word-break
package medium

func WordBreak(s string, wordDict []string) bool {
	return bottomUpWordBreak(s, wordDict)
	// return topDownWordBreak(s, wordDict)
}

func bottomUpWordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make(map[int]bool, n+1)
	dp[n] = true

	for i := n - 1; i >= 0; i-- {
		for _, w := range wordDict {
			if (i+len(w) <= n) && s[i:i+len(w)] == w {
				dp[i] = dp[i+len(w)]
			}
			if dp[i] {
				break
			}
		}
	}
	return dp[0]
}

func topDownWordBreak(s string, wordDict []string) bool {
	memo := make([]int, len(s))
	for i := range memo {
		memo[i] = -1
	}
	return isValid(s, wordDict, 0, memo)
}

func isValid(s string, wordDict []string, idx int, memo []int) bool {
	if idx == len(s) {
		return true
	}
	if memo[idx] != -1 {
		return memo[idx] == 1
	}

	for _, word := range wordDict {
		wordLen := len(word)
		if idx+wordLen > len(s) {
			continue
		}
		if s[idx:idx+wordLen] == word {
			if isValid(s, wordDict, idx+wordLen, memo) {
				memo[idx] = 1
				return true
			}
		}
	}

	memo[idx] = 0
	return false
}
