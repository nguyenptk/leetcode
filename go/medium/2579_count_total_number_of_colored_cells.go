// https://leetcode.com/problems/count-total-number-of-colored-cells/
package medium

func ColoredCells(n int) int64 {
	result := int64(1 + 4*n*(n-1)/2)
	return result
}
