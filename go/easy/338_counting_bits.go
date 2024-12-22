// https://leetcode.com/problems/counting-bits/
package easy

func CountBits(n int) []int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		result[i] = popCount(i)
	}

	return result
}

func popCount(idx int) int {
	count := 0
	for count = 0; idx != 0; count++ {
		idx = idx & (idx - 1)
	}

	return count
}