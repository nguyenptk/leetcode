// https://leetcode.com/problems/longest-palindromic-substring/
package medium

func LongestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	start := 0
	maxLen := 1
	var low, high int

	for i := 0; i < n; i++ {
		low = i - 1
		high = i + 1
		for high < n && s[high] == s[i] {
			high++
		}
		for low >= 0 && s[low] == s[i] {
			low--
		}
		for high < n && low >= 0 && s[high] == s[low] {
			low--
			high++
		}
		len := high - low - 1
		if maxLen < len {
			maxLen = len
			start = low + 1
		}
	}

	return s[start : start+maxLen]
}
