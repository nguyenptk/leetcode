// https://leetcode.com/problems/push-dominoes/
package medium

func PushDominoes(dominoes string) string {
	dominoesBytes := []byte(dominoes)

	l := -1
	r := -1

	for i := 0; i <= len(dominoes); i++ {
		if i == len(dominoes) || dominoes[i] == 'R' {
			if l < r {
				for r < i {
					dominoesBytes[r] = 'R'
					r++
				}
			}
			r = i
		} else if dominoes[i] == 'L' {
			if r < l || l == -1 && r == -1 {
				if l == -1 && r == -1 {
					l++
				}
				for l < i {
					dominoesBytes[l] = 'L'
					l++
				}

			} else {
				l = r + 1
				r = i - 1
				for l < r {
					dominoesBytes[l] = 'R'
					dominoesBytes[r] = 'L'
					l++
					r--
				}
			}
			l = i
		}
	}
	return string(dominoesBytes)
}
