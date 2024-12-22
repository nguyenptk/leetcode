// https://leetcode.com/problems/sum-of-two-integers/
package medium

func getSum(a int, b int) int {
	for b != 0 {
		answer := a ^ b
		carry := (a & b) << 1
		a = answer
		b = carry
	}
	return a
}
