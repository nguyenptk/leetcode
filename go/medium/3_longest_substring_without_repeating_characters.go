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

// Time complexity: O(2n)=O(n)
// Space complexity: O(min(m,n))

// Optimized approach
// func lengthOfLongestSubstring(s string) int {
// 	freq := make(map[byte]int)
// 	maxLen := 0
// 	l := 0
// 	for r := 0; r < len(s); r++ {
// 		if prev, ok := freq[s[r]]; ok {
// 			l = max(l, prev)
// 		}
// 		maxLen = max(maxLen, r-l+1)
// 		freq[s[r]] = r + 1
// 	}
// 	return maxLen
// }
// Time complexity: O(n)
// Space complexity: O(m)
