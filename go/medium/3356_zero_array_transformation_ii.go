// https://leetcode.com/problems/zero-array-transformation-ii/
package medium

func MinZeroArray(nums []int, queries [][]int) int {
	n := len(nums)
	sum := 0
	k := 0
	diff := make([]int, n+1)

	for i := 0; i < n; i++ {
		for sum+diff[i] < nums[i] {
			k++

			if k > len(queries) {
				return -1
			}
			l := queries[k-1][0]
			r := queries[k-1][1]
			val := queries[k-1][2]

			if r >= i {
				diff[max(l, i)] += val
				diff[r+1] -= val
			}
		}
		sum += diff[i]
	}

	return k
}
