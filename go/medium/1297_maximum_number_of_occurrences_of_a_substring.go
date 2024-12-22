// https://leetcode.com/problems/maximum-number-of-occurrences-of-a-substring/
package medium

func MaxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	result := 0
	freqs := make(map[string]int)

	for i := 0; i < len(s)-minSize+1; i++ {
		sub := s[i : i+minSize]
		unique := make(map[byte]bool)
		for j := 0; j < len(sub); j++ {
			unique[sub[j]] = true
		}
		if len(unique) <= maxLetters {
			freqs[sub]++
			if result < freqs[sub] {
				result = freqs[sub]
			}
		}
	}

	return result
}
