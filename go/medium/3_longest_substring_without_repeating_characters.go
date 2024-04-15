// https://leetcode.com/problems/longest-substring-without-repeating-characters/
package medium

func LengthOfLongestSubstring(s string) int {
	chars := map[byte]bool{}

	l := 0
	result := 0

	for r := 0; r <= len(s)-1; r++ {
		// check duplicated
		// look the chars with s[r] element
		for chars[s[r]] {
			delete(chars, s[l])
			l++
		}
		chars[s[r]] = true
		result = max(result, r-l+1)
	}

	return result
}
