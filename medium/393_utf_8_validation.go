// https://leetcode.com/problems/utf-8-validation/
package medium

func ValidUtf8(data []int) bool {
	leftToCheck := 0

	for _, v := range data {
		if leftToCheck == 0 {
			if (v >> 3) == 0b11110 {
				leftToCheck = 3
			} else if (v >> 4) == 0b1110 {
				leftToCheck = 2
			} else if (v >> 5) == 0b110 {
				leftToCheck = 1
			} else if (v >> 7) == 0b0 {
				leftToCheck = 0
			} else {
				return false
			}
		} else {
			if (v >> 6) != 0b10 {
				return false
			}
			leftToCheck--
		}
	}

	return leftToCheck == 0
}
