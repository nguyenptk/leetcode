// https://leetcode.com/problems/longest-valid-parentheses/
package hard

func LongestValidParentheses(s string) int {
	result := 0
	stack := make([]int, 0)
	stack = append(stack, -1)

	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				result = max(result, i-stack[len(stack)-1])
			}
		} else {
			stack = append(stack, i)
		}

	}

	return result
}
