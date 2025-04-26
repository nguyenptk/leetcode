// https://leetcode.com/problems/longest-substring-with-at-most-two-distinct-characters/
package medium

func LengthOfLongestSubstringTwoDistinct(s string) int {
	result := 0
	m := map[byte]int{}
	l := 0

	for r := 0; r < len(s); r++ {
		m[s[r]]++
		for len(m) > 2 {
			m[s[l]]--
			if m[s[l]] == 0 {
				delete(m, s[l])
			}
			l++
		}

		result = max(result, r-l+1)
	}

	return result
}
