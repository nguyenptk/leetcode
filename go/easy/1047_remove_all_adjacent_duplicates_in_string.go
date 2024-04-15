// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/
package easy

func RemoveAllAdjacentDuplicates(s string) string {
	i := 0
	for i < len(s)-1 {
		if s[i] == s[i+1] {
			tmp := s[:i] + s[i+2:]
			s = tmp
			i = 0
		} else {
			i++
		}
	}
	return s
}
