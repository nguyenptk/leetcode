// https://leetcode.com/problems/longest-substring-without-repeating-characters/
package medium

func LengthOfLongestSubstring(s string) int {
	v := map[rune]int{}
	max := 0
	p := 0
	n := 0
	for i, r := range s {
		if v[r] > p {
			p = v[r]
		}
		n = i - p + 1
		if max < n {
			max = n
		}
		v[r] = i + 1
	}
	return max
}
