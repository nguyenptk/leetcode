// https://leetcode.com/problems/check-if-a-parentheses-string-can-be-valid/
package medium

func CanBeValid(s string, locked string) bool {
	n := len(s)
	if n%2 != 0 {
		return false
	}

	maxOpen := 0
	minOpen := 0

	for i := 0; i < n; i++ {
		if locked[i] == '0' {
			maxOpen++
			if minOpen > 0 {
				minOpen--
			}
		} else {
			if s[i] == '(' {
				maxOpen++
				minOpen++
			} else {
				maxOpen--
				if minOpen > 0 {
					minOpen--
				}
				if minOpen > maxOpen {
					return false
				}
			}
		}
	}

	return minOpen == 0
}
