// https://leetcode.com/problems/647/
package medium

func CountSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isPalindromic(s, i, j) {
				count++
			}
		}
	}
	return count
}

func isPalindromic(s string, i, j int) bool {
	if i > j {
		return true
	}
	if s[i] != s[j] {
		return false
	}
	return isPalindromic(s, i+1, j-1)
}
