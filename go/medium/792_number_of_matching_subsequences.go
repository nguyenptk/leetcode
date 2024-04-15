// https://leetcode.com/problems/number-of-matching-subsequences/
package medium

func NumMatchingSubseq(s string, words []string) int {
	result := 0
	for _, w := range words {
		if len(words) > len(s) {
			continue
		}
		if isSubsq(s, w) {
			result++
		}
	}
	return result
}

func isSubsq(s, w string) bool {
	wIndex := 0
	for i := 0; i < len(s) && wIndex < len(w); i++ {
		if s[i] == w[wIndex] {
			wIndex++
		}
	}
	return wIndex == len(w)
}
