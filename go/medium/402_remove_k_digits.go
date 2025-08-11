// https://leetcode.com/problems/remove-k-digits/
package medium

func RemoveKdigits(num string, k int) string {
	stack := make([]byte, 0)
	for i := range num {
		digit := num[i]
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > digit {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
	}

	// If k > 0, remove remaining digits from the end
	stack = stack[:len(stack)-k]

	start := 0
	for start < len(stack) && stack[start] == '0' {
		start++
	}

	if start == len(stack) {
		return "0"
	}

	return string(stack[start:])
}

// Time complexity : O(N)
// Together with the outer loop, we have the exact (N+k) number of operations. Since 0<k≤N, the time complexity of the main loop is bounded within 2N.
// Space complexity : O(N)
