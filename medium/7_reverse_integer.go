// https://leetcode.com/problems/reverse-integer/
package medium

func Reverse(x int) int {
	y := x
	carry := 0
	for y/10 != 0 {
		carry = (carry * 10) + (y % 10)
		y = y / 10
	}
	if x < 0 {
		if 0^(carry*10+y) < -2147483648 {
			return 0
		}
		return 0 ^ (carry*10 + y)
	}
	if (carry*10 + y) > 2147483647 {
		return 0
	}
	return (carry*10 + y)
}
