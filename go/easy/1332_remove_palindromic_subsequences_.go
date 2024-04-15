// https://leetcode.com/problems/remove-palindromic-subsequences/
package easy

func RemovePalindromeSub(s string) int {
	if s == "" {
		return 0
	}

	flag := true

	start := 0
	last := len(s) - 1
	for last > start {
		if s[start] != s[last] {
			flag = false
		}
		start++
		last--
	}

	if !flag {
		return 2
	}

	return 1
}
