// https://leetcode.com/problems/remove-invalid-parentheses/
package hard

func RemoveInvalidParentheses(s string) []string {
	l := 0
	r := 0
	for i := range s {
		if s[i] == '(' {
			l++
		} else if s[i] == ')' {
			if l == 0 {
				r++
			}
			if l > 0 {
				l--
			}
		}
	}

	resultSet := make(map[string]bool)

	var backtrack func(idx, lCount, rCount, lRemain, rRemain int, path []byte)
	backtrack = func(idx, lCount, rCount, lRemain, rRemain int, path []byte) {
		if idx == len(s) {
			if lRemain == 0 && rRemain == 0 {
				resultSet[string(path)] = true
			}
			return
		} else {
			// Option to remove '('
			if s[idx] == '(' && lRemain > 0 {
				backtrack(idx+1, lCount, rCount, lRemain-1, rRemain, path)
			}
			// Option to remove ')'
			if s[idx] == ')' && rRemain > 0 {
				backtrack(idx+1, lCount, rCount, lRemain, rRemain-1, path)
			}
			path = append(path, s[idx])

			// Simply recurse one step further if the current character is not a parenthesis.
			if s[idx] != '(' && s[idx] != ')' {
				backtrack(idx+1, lCount, rCount, lRemain, rRemain, path)
			} else if s[idx] == '(' {
				// Consider an opening bracket.
				backtrack(idx+1, lCount+1, rCount, lRemain, rRemain, path)
			} else if s[idx] == ')' && lCount > rCount {
				// Consider an closing bracket.
				backtrack(idx+1, lCount, rCount+1, lRemain, rRemain, path)
			}

			path = path[:len(path)-1]
		}
	}

	backtrack(0, 0, 0, l, r, []byte{})

	result := make([]string, 0, len(resultSet))
	for str := range resultSet {
		result = append(result, str)
	}
	return result
}

// Time complexity: in the worst case we still have 2 options per parenthesis and that gives us a complexity of O(2^N).
// Space complexity: O(N)
