// https://leetcode.com/problems/kth-missing-positive-number/
package easy

func FindKthPositive(arr []int, k int) int {
	// original = [1,2,3,4,5,6,7,8,9,10]
	// arr      = [2,3,4,7,11]
	l, r := 0, len(arr)-1
	for l < r {
		m := (l + r) / 2
		if arr[m]-(m+1) < k {
			l = m + 1
		} else {
			r = m
		}
	}
	return l + k
}

// Time complexity: log(N)
// Space complexity: N
