// https://leetcode.com/problems/longest-repeating-character-replacement
package medium

func CharacterReplacement(s string, k int) int {
	count := map[byte]int{}

	result := 0
	l := 0
	maxF := 0
	for r := 0; r < len(s); r++ {
		count[s[r]] += 1
		maxF = max(maxF, count[s[r]])
		for (r-l+1)-maxF > k {
			count[s[l]] -= 1
			l++
		}
		result = max(result, r-l+1)
	}

	return result
}
