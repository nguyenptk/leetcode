// https://leetcode.com/problems/spiral-matrix/
package medium

func SpiralOrder(matrix [][]int) []int {
	rows := len(matrix)
	cols := len(matrix[0])
	result := make([]int, 0)

	top := 0
	left := 0
	right := cols
	bottom := rows

	for left < right && top < bottom {
		for i := left; i < right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		for i := top; i < bottom; i++ {
			result = append(result, matrix[i][right-1])
		}
		right--

		if left == right || top == bottom {
			break
		}

		for i := right - 1; i >= left; i-- {
			result = append(result, matrix[bottom-1][i])
		}
		bottom--

		for i := bottom - 1; i >= top; i-- {
			result = append(result, matrix[i][left])
		}
		left++

	}

	return result
}
