// https://leetcode.com/problems/permutation-in-string
package medium

func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	m := [26]int{}

	for i := range s1 {
		m[s1[i]-'a']++
	}

	length := len(s1)
	l := 0

	for r := 0; r < len(s2); r++ {
		if m[s2[r]-'a'] > 0 {
			length--
		}
		m[s2[r]-'a']--

		if length == 0 {
			return true
		}
		if r >= len(s1)-1 {
			if m[s2[l]-'a'] >= 0 {
				length++
			}
			m[s2[l]-'a']++
			l++
		}
	}

	return false
}
