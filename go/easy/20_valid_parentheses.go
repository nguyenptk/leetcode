// https://leetcode.com/problems/valid-parentheses/
package easy

func IsValid(s string) bool {
	stack := []rune{}
	closeToOpen := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, c := range s {
		// Check the current char is the closing bracket
		if v, ok := closeToOpen[c]; ok {
			top := '#'
			if len(stack) > 0 {
				top = stack[len(stack)-1]    // Top
				stack = stack[:len(stack)-1] // Pop
			}
			if v != top {
				return false
			}
		} else {
			stack = append(stack, c)
		}
	}

	return len(stack) == 0
}
