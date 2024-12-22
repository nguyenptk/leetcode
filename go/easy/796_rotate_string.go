// https://leetcode.com/problems/rotate-string/

package easy

func RotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	for i := 0; i <= len(s); i++ {
		first := s[:1]
		others := s[1:]
		s = others + first
		if s == goal {
			return true
		}
	}

	return false
}
