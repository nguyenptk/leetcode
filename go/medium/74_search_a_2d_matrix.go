// https://leetcode.com/problems/search-a-2d-matrix
package medium

func SearchMatrixI(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	l := 0
	r := m*n - 1

	for l <= r {
		mIdx := (l + r) / 2            // middle index of the array
		mEle := matrix[mIdx/n][mIdx%n] // middle element of the array
		if target == mEle {
			return true
		} else if target > mEle {
			l = mIdx + 1
		} else {
			r = mIdx - 1
		}
	}
	return false
}
