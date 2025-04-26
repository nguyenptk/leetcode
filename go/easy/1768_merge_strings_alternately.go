// https://leetcode.com/problems/merge-strings-alternately/
package easy

func MergeAlternately(word1 string, word2 string) string {
	minL := len(word1)
	if len(word1) > len(word2) {
		minL = len(word2)
	}

	result := make([]byte, 0)
	i := 0
	for ; i < minL; i++ {
		if i == len(word1) || i == len(word2) {
			break
		}
		result = append(result, word1[i])
		result = append(result, word2[i])
	}

	for i < len(word1) || i < len(word2) {
		if i < len(word1) {
			result = append(result, word1[i])
		} else if i < len(word2) {
			result = append(result, word2[i])
		}
		i++
	}

	return string(result)
}
