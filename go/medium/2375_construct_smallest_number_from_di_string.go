// https://leetcode.com/problems/construct-smallest-number-from-di-string/
package medium

import "strconv"

func SmallestNumber(pattern string) string {
	result := ""

	stack := make([]int, 0)
	for i := 0; i <= len(pattern); i++ {
		stack = append(stack, i+1)
		if i == len(pattern) || pattern[i] == 'I' {
			for len(stack) > 0 {
				result += strconv.Itoa(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
		}
	}

	return result
}
