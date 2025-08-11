// https://leetcode.com/problems/score-of-parentheses/
package medium

func ScoreOfParentheses(s string) int {
	result := 0
	bal := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			bal++
		} else {
			bal--
			if s[i-1] == '(' {
				result += 1 << bal // 2^bal
			}
		}
	}
	return result
}
