// https://leetcode.com/problems/longest-valid-parentheses/
package hard

func LongestValidParentheses(s string) int {
	// return early if it's not enough length
	if len(s) < 2 {
		return 0
	}

	// init result
	result := 0

	// init rune stack and push initital index
	stack := []rune{}
	stack = append(stack, -1)

	// loop the give string
	for i, v := range s {
		// if opening bracket, push index of it
		if string(v) == "(" {
			stack = append(stack, rune(i))

			// if closing bracket
		} else {
			// pop the previous opening bracket's index
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}

			// check if this length formed with base of
			// current valid substring is more than max
			// so far
			if len(stack) != 0 {
				n := int(stack[len(stack)-1])
				if i-n > result {
					result = i - n
				}

				// if stack is empty. push current index as
				// base for next valid substring (if any)
			} else {
				stack = append(stack, rune(i))
			}
		}
	}
	return result
}
