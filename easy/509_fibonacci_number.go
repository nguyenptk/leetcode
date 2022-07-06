// https://leetcode.com/problems/fibonacci-number/
package easy

func Fib(n int) int {
	if n <= 1 {
		return n
	}
	return Fib(n-2) + Fib(n-1)
}
