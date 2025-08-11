// https://leetcode.com/problems/k-th-smallest-in-lexicographical-order
package hard

func FindKthNumber(n int, k int) int {
	curr := 1
	k--

	var countSteps func(prefix1, prefix2 int) int
	countSteps = func(prefix1, prefix2 int) int {
		steps := 0
		for prefix1 <= n {
			steps += min(n+1, prefix2) - prefix1
			prefix1 *= 10
			prefix2 *= 10
		}
		return steps
	}

	for k > 0 {
		step := countSteps(curr, curr+1)
		if step <= k {
			curr++
			k -= step
		} else {
			curr *= 10
			k--
		}
	}

	return curr
}

// Time complexity: O(log(n)^2)
// The outer while loop runs as long as k > 0. In the worst case, it runs O(logn) times, because at each step, we either move to the next prefix or move deeper into the tree (multiplying the current prefix by 10).
// The countSteps function, which calculates the number of steps between two prefixes, runs in O(logn) time, as it traverses deeper levels of the number range by multiplying the prefixes by 10 in each iteration.
// Since the countSteps function is called inside the while loop, which also runs O(logn) times, the overall time complexity is O(log(n)×log(n))=O(log(n)^2).
// Space complexity: O(1)
// The space complexity is O(1) because we're only using a constant amount of additional space for variables like curr, k, step, prefix1, prefix2, etc. We're not using any data structures that grow with the input size.
