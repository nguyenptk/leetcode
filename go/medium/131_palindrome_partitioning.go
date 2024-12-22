// https://leetcode.com/problems/palindrome-partitioning/

package medium

func PalindromePartition(s string) [][]string {
	result := make([][]string, 0)
	curr := make([]string, 0)
	var backtrack func(idx int)
	backtrack = func(idx int) {
		if idx == len(s) {
			result = append(result, append([]string{}, curr...))
		}
		for i := idx; i < len(s); i++ {
			if isPalindrome(s[idx : i+1]) {
				curr = append(curr, s[idx:i+1])
				backtrack(i + 1)
				curr = curr[:len(curr)-1]
			}
		}
	}
	backtrack(0)

	return result
}

func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
