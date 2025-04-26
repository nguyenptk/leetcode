// https://leetcode.com/problems/basic-calculator-iii/
package hard

func CalculateIII(s string) int {
	s += "0" // treat the end‑of‑string like a “0” sentinel
	idx := 0

	var eval func() int
	eval = func() int {
		result, lastNum, curNum := 0, 0, 0
		operator := byte('+')

		for idx < len(s) {
			ch := s[idx]
			idx++

			if ch >= '0' && ch <= '9' {
				curNum = curNum*10 + int(ch-'0')
			}

			if ch == '(' {
				curNum = eval()
			}

			// if this is an operator, a ')', or we're at the end, apply the previous sign
			if (ch < '0' || ch > '9') || idx == len(s)-1 {
				switch operator {
				case '+':
					result += lastNum
					lastNum = curNum
				case '-':
					result += lastNum
					lastNum = -curNum
				case '*':
					lastNum *= curNum
				case '/':
					lastNum /= curNum
				}

				if ch == ')' {
					break
				}
				operator = ch
				curNum = 0
			}
		}

		return result + lastNum
	}

	return eval()
}
