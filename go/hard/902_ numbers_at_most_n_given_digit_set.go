// https://leetcode.com/problems/numbers-at-most-n-given-digit-set/
package hard

import (
	"math"
	"strconv"
)

func AtMostNGivenDigitSet(digits []string, n int) int {
	nStr := strconv.Itoa(n)
	nLen := len(nStr)
	digitLen := len(digits)

	result := 0

	// Count numbers with length less than n
	for i := 1; i < nLen; i++ {
		result += int(math.Pow(float64(digitLen), float64(i)))
	}

	// count numbers with the same length
	for i := 0; i < nLen; i++ {
		hasSameDigit := false
		for _, d := range digits {
			if d[0] < nStr[i] {
				result += int(math.Pow(float64(digitLen), float64(nLen-i-1)))
			} else if d[0] == nStr[i] {
				hasSameDigit = true
				break
			}
		}
		if !hasSameDigit {
			return result
		}
	}

	return result + 1
}
