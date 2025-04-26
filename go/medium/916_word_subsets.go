// https://leetcode.com/problems/word-subsets/
package medium

func WordSubsets(words1 []string, words2 []string) []string {
	result := make([]string, 0)

	maxFreq := make(map[byte]int)
	for _, word2 := range words2 {
		tempFreq := make(map[byte]int)
		for i := 0; i < len(word2); i++ {
			tempFreq[word2[i]]++
			maxFreq[word2[i]] = max(maxFreq[word2[i]], tempFreq[word2[i]])
		}
	}

	for _, word1 := range words1 {
		m := make(map[byte]int)
		for i := 0; i < len(word1); i++ {
			m[word1[i]]++
		}

		isSubset := true
		for c, count := range maxFreq {
			if m[c] < count {
				isSubset = false
				break
			}
		}

		if isSubset {
			result = append(result, word1)
		}
	}

	return result
}
