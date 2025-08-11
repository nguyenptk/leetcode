// https://leetcode.com/problems/number-of-segments-in-a-string/
package easy

func CountSegments(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		if (i == 0 || s[i-1] == ' ') && s[i] != ' ' {
			result++
		}
	}
	return result
}
