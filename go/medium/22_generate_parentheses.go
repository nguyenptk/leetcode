// https://leetcode.com/problems/generate-parentheses
package medium

func GenerateParenthesis(n int) []string {
	result := make([]string, 0)

	var backtrack func(open, close int, s string)
	backtrack = func(open, close int, s string) {
		if len(s) == n*2 {
			result = append(result, s)
			s = ""
			return
		}

		if open < n {
			s += "("
			backtrack(open+1, close, s)
			s = s[:len(s)-1]
		}
		if open > close {
			s += ")"
			backtrack(open, close+1, s)
			s = s[:len(s)-1]
		}

	}
	backtrack(0, 0, "")

	return result
}
