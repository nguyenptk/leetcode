// https://leetcode.com/problems/length-of-last-word/
package easy

func LengthOfLastWord(s string) int {
	result := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if result == 0 {
				continue
			} else {
				return result
			}
		}
		result++
	}

	return result
}
