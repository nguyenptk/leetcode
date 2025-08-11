// https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/
package easy

func MaxDepth(s string) int {
	result := 0
	depth := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			depth++
			result = max(result, depth)
		} else if s[i] == ')' {
			depth--
		}
	}

	return result
}
