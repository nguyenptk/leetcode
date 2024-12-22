// https://leetcode.com/problems/longest-palindromic-substring/
package medium

func LongestPalindrome(s string) string {
	n := len(s)
	result := ""
	maxLen := 0

	var getLongest func(int, int)
	getLongest = func(l, r int) {
		for l >= 0 && r < n && s[l] == s[r] {
			if r-l+1 > maxLen {
				maxLen = r - l + 1
				result = s[l : r+1]
			}
			l--
			r++
		}
	}

	for i := 0; i < n; i++ {
		// odd length
		getLongest(i, i)

		// even length
		getLongest(i, i+1)
	}

	return result
}
