// https://leetcode.com/problems/add-strings/
package easy

import (
	"strconv"
)

func AddStrings(num1 string, num2 string) string {
	p1 := len(num1) - 1
	p2 := len(num2) - 1

	sum := ""
	carry := 0
	for p1 >= 0 || p2 >= 0 {
		x1 := 0
		if p1 >= 0 {
			x1 = int(num1[p1] - '0')
		}
		x2 := 0
		if p2 >= 0 {
			x2 = int(num2[p2] - '0')
		}
		val := (x1 + x2 + carry) % 10
		carry = (x1 + x2 + carry) / 10
		sum = sum + strconv.Itoa(val)
		p1--
		p2--
	}

	if carry != 0 {
		sum = sum + strconv.Itoa(carry)
	}

	result := ""
	for i := len(sum) - 1; i >= 0; i-- {
		result += string(sum[i])
	}

	return result
}
