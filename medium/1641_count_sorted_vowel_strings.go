// https://leetcode.com/problems/count-sorted-vowel-strings/

package medium

func CountVowelStrings(n int) int {
	return countVowels(n, 0)
}

// Recursive count vowels
func countVowels(n int, index int) int {
	if n == 0 {
		return 1
	}
	count := 0
	// Loop the vowels ["a", "e", "i", "o", "u"]
	for i := index; i < 5; i++ {
		// Recursive to plus the count
		count += countVowels(n-1, i)
	}
	return count
}
