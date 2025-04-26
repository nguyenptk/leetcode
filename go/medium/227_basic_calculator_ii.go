// https://leetcode.com/problems/basic-calculator-ii/
package medium

func CalculateII(s string) int {
	result := 0
	operator := byte('+')
	last := 0
	curr := 0

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			curr = curr*10 + int(s[i]-'0')
		}
		if (s[i] < '0' || s[i] > '9') && s[i] != ' ' || i == len(s)-1 {
			if operator == '+' {
				result += last
				last = curr
			} else if operator == '-' {
				result += last
				last = -curr
			} else if operator == '*' {
				last = last * curr
			} else if operator == '/' {
				last = last / curr
			}
			operator = s[i]
			curr = 0
		}
	}

	return result + last
}
