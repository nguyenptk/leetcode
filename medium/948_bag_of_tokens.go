// https://leetcode.com/problems/bag-of-tokens/
package medium

import "sort"

func BagOfTokensScore(tokens []int, power int) int {
	result := 0
	score := 0
	i := 0               // index of smallest token
	j := len(tokens) - 1 // index of largest token

	sort.Slice(tokens, func(i, j int) bool {
		return tokens[i] < tokens[j]
	})

	for i <= j && (power >= tokens[i] || score > 0) {
		for i <= j && power >= tokens[i] {
			// play the smallest face up
			power -= tokens[i]
			i++
			score++
		}
		if result < score {
			result = score
		}
		if i <= j && score > 0 {
			// play the largest face down
			power += tokens[j]
			j--
			score--
		}
	}

	return result
}
