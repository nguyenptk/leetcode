// https://leetcode.com/problems/decode-string/
package medium

import (
	"strings"
	"unicode"
)

func DecodeString(s string) string {
	i := 0
	var rec func() string
	rec = func() string {
		var result strings.Builder
		for i < len(s) && s[i] != ']' {
			if !unicode.IsDigit(rune(s[i])) {
				result.WriteByte(s[i])
				i++
			} else {
				k := 0
				// build k while next character is a digit
				for i < len(s) && unicode.IsDigit(rune(s[i])) {
					k = k*10 + int(s[i]-'0')
					i++
				}
				// ignore the opening bracket '['
				i++
				decode := rec()
				// ignore the closing bracket ']'
				i++
				for k > 0 {
					result.WriteString(decode)
					k--
				}
			}
		}
		return result.String()
	}

	return rec()
}

// Time Complexity: O(maxK⋅n) where maxK is the maximum value of k and n is the length of a given string s.
// Space Complexity: O(n).

// func DecodeString(s string) string {
// 	stack := make([]interface{}, 0)

// 	curStr := ""
// 	curNum := 0

// 	for i := 0; i < len(s); i++ {
// 		if s[i] == '[' {
// 			stack = append(stack, curStr)
// 			stack = append(stack, curNum)
// 			curStr = ""
// 			curNum = 0

// 		} else if s[i] == ']' {
// 			num := stack[len(stack)-1]
// 			stack = stack[:len(stack)-1]

// 			str := stack[len(stack)-1]
// 			stack = stack[:len(stack)-1]

// 			curStr = str.(string) + strings.Repeat(curStr, num.(int))

// 		} else if s[i] >= '0' && s[i] <= '9' {
// 			digit := s[i] - '0'
// 			curNum = curNum*10 + int(digit)

// 		} else {
// 			curStr += string(s[i])
// 		}
// 	}

// 	return curStr
// }
