// https://leetcode.com/problems/valid-anagram/
package easy

import (
	"sort"
)

func IsAnagram(s string, t string) bool {
	// basic case
	if len(s) != len(t) {
		return false
	}
	// sort s & t
	sArr := []rune(s)
	sort.Slice(sArr, func(i, j int) bool {
		return sArr[i] < sArr[j]
	})
	tArr := []rune(t)
	sort.Slice([]rune(tArr), func(i, j int) bool {
		return tArr[i] < tArr[j]
	})
	// compare
	for i := 0; i < len(sArr); i++ {
		if sArr[i] != tArr[i] {
			return false
		}
	}
	return true
}
