// https://leetcode.com/problems/fibonacci-number/
package easy

func Fib(n int) int {
	if n <= 1 {
		return n
	}
	memo := make([]int, n+1)
	memo[1] = 1

	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}

	return memo[n]
}
