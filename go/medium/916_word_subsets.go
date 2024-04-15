// https://leetcode.com/problems/word-subsets/
package medium

func WordSubsets(words1 []string, words2 []string) []string {
	result := []string{}
	counts2 := make([]int, 26)

	for _, w2 := range words2 {
		tmp := counter(w2)
		for i := 0; i < 26; i++ {
			if counts2[i] < tmp[i] {
				counts2[i] = tmp[i]
			}
		}
	}

	for _, w1 := range words1 {
		if isUniversal(counter(w1), counts2) {
			result = append(result, w1)
		}

	}

	return result
}

func counter(s string) []int {
	count := make([]int, 26)
	for _, c := range s {
		count[c-'a']++
	}
	return count
}

func isUniversal(counts1 []int, counts2 []int) bool {
	for i := 0; i < 26; i++ {
		if counts1[i] < counts2[i] {
			return false
		}
	}
	return true
}
