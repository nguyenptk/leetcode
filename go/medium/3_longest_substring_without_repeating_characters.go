// https://leetcode.com/problems/longest-substring-without-repeating-characters/
package medium

func LengthOfLongestSubstring(s string) int {
	result := 0
	freq := make([]bool, 128)
	l := 0
	for r := 0; r < len(s); r++ {
		for freq[s[r]] {
			freq[s[l]] = false
			l++
		}
		freq[s[r]] = true
		result = max(result, r-l+1)
	}

	return result
}
