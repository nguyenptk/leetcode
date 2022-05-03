// https://leetcode.com/problems/valid-parentheses/
package easy

import "strings"

func IsValid(s string) bool {
	// return early if it is not enough length
	if len(s) < 2 {
		return false
	}

	// remove all space
	s = strings.ReplaceAll(s, " ", "")

	// check even length
	if len(s)%2 != 0 {
		return false
	}

	// init rune stack and openClose map
	stack := []rune{}
	openClose := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, v := range s {
		// if there is "(", "[" or "{" character, the ")", "]" or "}" character will be added
		if closer, ok := openClose[v]; ok {
			stack = append(stack, closer)
			continue
		}
		// check stack is not empty and whether the top of stack matches the current character
		l := len(stack) - 1
		if l < 0 || v != stack[l] {
			return false
		}
		stack = stack[:l]
	}

	return len(stack) == 0
}
