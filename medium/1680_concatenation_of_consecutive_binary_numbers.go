// https://leetcode.com/problems/concatenation-of-consecutive-binary-numbers/
package medium

func ConcatenatedBinary(n int) int {
	result := 1
	len := 4
	for i := 2; i <= n; i++ {
		if i == len {
			len *= 2
		}
		result = (result*len + i) % 1000000007
	}
	return result
}
