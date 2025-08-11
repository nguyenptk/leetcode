// https://leetcode.com/problems/powx-n/
package medium

func MyPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / MyPow(x, -1*n)
	}
	// Perform Binary Exponentiation.
	// If 'n' is odd we perform Binary Exponentiation on 'n - 1' and multiply result with 'x'.
	if n%2 == 1 {
		return x * MyPow(x*x, (n-1)/2)
	}
	// Otherwise we calculate result by performing Binary Exponentiation on 'n'.
	return MyPow(x*x, n/2)
}
