// https://leetcode.com/problems/longest-palindrome-by-concatenating-two-letter-words/solution/
package medium

func LongestPalindromeConcatenating(words []string) int {
	alphabetSize := 26
	count := make([][]int, alphabetSize)
	for i := range count {
		count[i] = make([]int, alphabetSize)
	}
	for _, w := range words {
		count[w[0]-'a'][w[1]-'a'] += 1
	}
	result := 0
	central := false
	for i := 0; i < alphabetSize; i++ {
		if count[i][i]%2 == 0 {
			result += count[i][i]
		} else {
			result += count[i][i] - 1
			central = true
		}
		for j := i + 1; j < alphabetSize; j++ {
			if count[i][j] < count[j][i] {
				result += 2 * count[i][j]
			} else {
				result += 2 * count[j][i]
			}
		}
	}
	if central {
		result++
	}
	return 2 * result
}
