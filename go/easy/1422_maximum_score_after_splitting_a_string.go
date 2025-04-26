// https://leetcode.com/problems/maximum-score-after-splitting-a-string/
package easy

func MaxScore(s string) int {
	// return bruteForceMaxScore(s)
	return countMaxScore(s)
}

func bruteForceMaxScore(s string) int {
	result := 0
	for i := 0; i < len(s)-1; i++ {
		curr := 0
		for j := 0; j <= i; j++ {
			if s[j] == '0' {
				curr++
			}
		}

		for j := i + 1; j < len(s); j++ {
			if s[j] == '1' {
				curr++
			}
		}
		result = max(result, curr)
	}

	return result
}

func countMaxScore(s string) int {
	result := 0

	one := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			one++
		}
	}

	zero := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			zero++
		} else {
			one--
		}
		result = max(result, zero+one)
	}

	return result
}
