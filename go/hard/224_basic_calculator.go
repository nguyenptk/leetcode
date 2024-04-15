// https://leetcode.com/problems/basic-calculator/
package hard

import "strconv"

func Calculate(s string) int {
	ans := 0
	num := 0
	sign := 1
	stack := []int{sign}

	for _, c := range s {
		if _, err := strconv.Atoi(string(c)); err == nil {
			num = num*10 + (int)(c-'0')
		} else if c == '(' {
			stack = append(stack, sign)
		} else if c == ')' {
			stack = stack[:len(stack)-1] // POP
		} else if c == '+' || c == '-' {
			ans += sign * num
			if c == '+' {
				sign = stack[len(stack)-1] // TOP
			} else {
				sign = -stack[len(stack)-1] // TOP
			}
			num = 0
		}
	}

	return ans + sign*num
}
