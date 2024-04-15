// https://leetcode.com/problems/make-the-string-great/
package easy

func MakeGood(s string) string {
	i := 0
	for i <= len(s)-2 {
		if s[i]-s[i+1] == 32 || s[i+1]-s[i] == 32 {
			tmp := s[:i] + s[i+2:]
			s = tmp
			i = 0
		} else {
			i++
		}
	}
	return s
}
