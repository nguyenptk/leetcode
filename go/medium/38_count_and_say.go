// https://leetcode.com/problems/count-and-say/
package medium

import "strconv"

func CountAndSay(n int) string {
	result := "1"

	for ; n > 1; n-- {
		next := ""
		i := 0
		for i < len(result) {
			count := 1
			for i < len(result)-1 && result[i] == result[i+1] {
				count++
				i++
			}
			next += strconv.Itoa(count) + string(result[i])
			i++
		}
		result = next
	}

	return result
}
