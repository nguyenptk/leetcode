// https://leetcode.com/problems/permutation-in-string
package medium

func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	// Init 2 arrays to store the frequency
	m1 := make([]int, 26)
	m2 := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		m1[s1[i]-'a'] += 1
		m2[s2[i]-'a'] += 1
	}

	matches := 0
	for i := 0; i < 26; i++ {
		if m1[i] == m2[i] {
			matches++
		}
	}

	l := 0
	for r := len(s1); r < len(s2); r++ {
		if matches == 26 {
			return true
		}

		index := s2[r] - 'a'
		m2[index]++
		if m1[index] == m2[index] {
			matches++
		} else if m1[index]+1 == m2[index] {
			matches--
		}

		index = s2[l] - 'a'
		m2[index]--
		if m1[index] == m2[index] {
			matches++
		} else if m1[index]-1 == m2[index] {
			matches--
		}
		l++
	}

	return matches == 26
}
