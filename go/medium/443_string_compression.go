// https://leetcode.com/problems/string-compression/
package medium

import (
	"fmt"
)

func Compress(chars []byte) int {
	count := 1
	l := 0

	for r := 1; r <= len(chars); r++ {
		if r < len(chars) && chars[r] == chars[r-1] {
			count++
		} else {
			chars[l] = chars[r-1]
			l++
			if count > 1 {
				for _, c := range []byte(fmt.Sprint(count)) {
					chars[l] = c
					l++
				}
			}
			count = 1
		}
	}

	return l
}
