// https://leetcode.com/problems/generate-parentheses
package medium

import "strings"

func GenerateParenthesis(n int) []string {
	result := []string{}
	stack := []string{}

	backtrack(&result, &stack, n, 0, 0)

	return result
}

func backtrack(result, stack *[]string, n, open, closed int) {
	if open == closed && open == n {
		*result = append(*result, strings.Join(*stack, ""))
	}

	if open < n {
		*stack = append(*stack, "(")
		backtrack(result, stack, n, open+1, closed)
		*stack = (*stack)[:len(*stack)-1]
	}

	if closed < open {
		*stack = append(*stack, ")")
		backtrack(result, stack, n, open, closed+1)
		*stack = (*stack)[:len(*stack)-1]
	}
}
