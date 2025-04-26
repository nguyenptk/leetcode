// https://leetcode.com/problems/shifting-letters-ii/
package medium

func ShiftingLetters(s string, shifts [][]int) string {
	n := len(s)
	result := make([]byte, n)

	prefixSum := make([]int, n)
	for _, shift := range shifts {
		start := shift[0]
		end := shift[1]
		direction := shift[2]
		if direction == 0 {
			prefixSum[start]--
			if end+1 < n {
				prefixSum[end+1]++
			}
		} else {
			prefixSum[start]++
			if end+1 < n {
				prefixSum[end+1]--
			}
		}
	}

	for i := 1; i < n; i++ {
		prefixSum[i] += prefixSum[i-1]
	}

	for i := 0; i < n; i++ {
		shift := prefixSum[i] % 26
		new := (int(s[i]-'a') + shift + 26) % 26
		result[i] = byte(new) + 'a'
	}

	return string(result)
}
