// https://leetcode.com/problems/transpose-matrix/
package easy

func Transpose(matrix [][]int) [][]int {
	xl := len(matrix[0])
	yl := len(matrix)
	result := make([][]int, xl)
	for i := range result {
		result[i] = make([]int, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = matrix[j][i]
		}
	}
	return result
}
