// https://leetcode.com/problems/valid-parenthesis-string/
package medium

func CheckValidString(s string) bool {
	n := len(s) - 1
	open := 0
	close := 0
	for i := 0; i <= n; i++ {
		if s[i] == '(' || s[i] == '*' {
			open++
		} else {
			open--
		}
		if s[n-i] == ')' || s[n-i] == '*' {
			close++
		} else {
			close--
		}
		if open < 0 || close < 0 {
			return false
		}
	}

	return true
}
