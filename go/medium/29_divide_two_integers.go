// https://leetcode.com/problems/divide-two-integers/
package medium

func Divide(dividend int, divisor int) int {
	if dividend/divisor == 2147483648 {
		return 2147483647
	}
	return dividend / divisor
}
