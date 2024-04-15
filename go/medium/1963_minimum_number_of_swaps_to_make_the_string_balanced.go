// https://leetcode.com/problems/minimum-number-of-swaps-to-make-the-string-balanced
package medium

func MinSwaps(s string) int {
	openClosed := 0
	maxS := 0
	for _, c := range s {
		if c == '[' {
			openClosed--
		} else if c == ']' {
			openClosed++
		}
		maxS = max(maxS, openClosed)
	}

	return (maxS + 1) / 2
}
