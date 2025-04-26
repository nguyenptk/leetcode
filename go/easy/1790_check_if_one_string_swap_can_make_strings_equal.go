// https://leetcode.com/problems/check-if-one-string-swap-can-make-strings-equal/
package easy

func areAlmostEqual(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	count := 0
	firstIdx := 0
	secondIdx := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			count++
			if count > 2 {
				return false
			} else if count == 1 {
				firstIdx = i
			} else {
				secondIdx = i
			}
		}
	}

	return s1[firstIdx] == s2[secondIdx] && s1[secondIdx] == s2[firstIdx]
}
