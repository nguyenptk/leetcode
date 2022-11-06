// https://leetcode.com/problems/orderly-queue/
package hard

import "sort"

func OrderlyQueue(s string, k int) string {
	if k == 1 {
		result := s
		for i := 0; i < len(s); i++ {
			tmp := s[i:] + s[0:i]
			if tmp < result {
				result = tmp
			}
		}
		return result
	} else {
		letters := []byte(s)
		sort.Slice(letters, func(i, j int) bool {
			return letters[i] < letters[j]
		})
		return string(letters)
	}
}
