// https://leetcode.com/problems/number-of-submatrices-that-currSum-to-target/
package hard

func NumSubmatrixSumTarget(matrix [][]int, target int) int {
	result := 0
	for r1 := 0; r1 < len(matrix); r1++ {
		sums := make([]int, len(matrix[0]))
		for r2 := r1; r2 < len(matrix); r2++ {
			m := map[int]int{
				0: 1,
			}
			rowSum := 0
			for x := range matrix[0] {
				rowSum += matrix[r2][x]
				sums[x] += rowSum
				result += m[sums[x]-target]
				m[sums[x]]++
			}
		}
	}

	return result
}
