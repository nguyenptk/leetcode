// https://leetcode.com/problems/decode-string/
package medium

import "strings"

func DecodeString(s string) string {
	stack := make([]interface{}, 0)

	curStr := ""
	curNum := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			stack = append(stack, curStr)
			stack = append(stack, curNum)
			curStr = ""
			curNum = 0

		} else if s[i] == ']' {
			num := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			str := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			curStr = str.(string) + strings.Repeat(curStr, num.(int))

		} else if s[i] >= '0' && s[i] <= '9' {
			digit := s[i] - '0'
			curNum = curNum*10 + int(digit)

		} else {
			curStr += string(s[i])
		}
	}

	return curStr
}
