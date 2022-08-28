// https://leetcode.com/problems/sort-the-matrix-diagonally/
package medium

import "sort"

func DiagonalSort(mat [][]int) [][]int {
	y := len(mat)
	x := len(mat[0]) - 1
	diag := make([]int, y)
	for i := 2 - y; i < x; i++ {
		k := 0
		var j int
		for j = 0; j < y; j++ {
			if i+j >= 0 && i+j <= x {
				diag[k] = mat[j][i+j]
				k++
			}
		}

		// sort 0 -> k elements in diag
		tmp := diag[0:k]
		sort.Slice(tmp, func(i, j int) bool {
			return diag[i] < diag[j]
		})
		diag = diag[:k]
		diag = append(tmp, diag...)

		k = 0
		for j = 0; j < y; j++ {
			if i+j >= 0 && i+j <= x {
				mat[j][i+j] = diag[k]
				k++
			}
		}
	}

	return mat
}
