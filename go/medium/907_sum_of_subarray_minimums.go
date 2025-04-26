// https://leetcode.com/problems/sum-of-subarray-minimums/
package medium

func SumSubarrayMins(arr []int) int {
	MOD := (int)(1e9 + 7)
	stack := make([]int, 0)
	result := 0
	for i := 0; i <= len(arr); i++ {
		for len(stack) > 0 && (i == len(arr) || arr[stack[len(stack)-1]] >= arr[i]) {
			m := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			l := -1
			if len(stack) > 0 {
				l = stack[len(stack)-1]
			}
			r := i

			count := (m - l) * (r - m)
			result += count * arr[m]
		}

		stack = append(stack, i)
	}

	return result % MOD
}
