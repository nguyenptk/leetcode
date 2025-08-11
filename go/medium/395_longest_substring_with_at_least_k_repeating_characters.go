// https://leetcode.com/problems/longest-substring-with-at-least-k-repeating-characters/
package medium

func LongestSubstring(s string, k int) int {
	maxU := maxUnique(s)
	result := 0

	for target := 1; target <= maxU; target++ {
		countMap := make([]int, 26)
		start, end := 0, 0
		unique, countAtLeastK := 0, 0

		for end < len(s) {
			if unique <= target {
				idx := s[end] - 'a'
				if countMap[idx] == 0 {
					unique++
				}
				countMap[idx]++
				if countMap[idx] == k {
					countAtLeastK++
				}
				end++
			} else {
				idx := s[start] - 'a'
				if countMap[idx] == k {
					countAtLeastK--
				}
				countMap[idx]--
				if countMap[idx] == 0 {
					unique--
				}
				start++
			}
			if unique == target && unique == countAtLeastK {
				result = max(result, end-start)
			}
		}
	}

	return result
}

func maxUnique(s string) int {
	seen := make([]bool, 26)
	cnt := 0
	for i := 0; i < len(s); i++ {
		idx := s[i] - 'a'
		if !seen[idx] {
			seen[idx] = true
			cnt++
		}
	}
	return cnt
}
