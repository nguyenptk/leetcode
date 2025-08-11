// https://leetcode.com/problems/sparse-matrix-multiplication/
package medium

func Multiply(mat1 [][]int, mat2 [][]int) [][]int {
	m, k := len(mat1), len(mat1[0])
	n := len(mat2[0])

	// Step 1: Compress mat1 — row-wise non-zero elements
	mMat1 := make([]map[int]int, m)
	for i := 0; i < m; i++ {
		mMat1[i] = make(map[int]int)
		for j := 0; j < k; j++ {
			if mat1[i][j] != 0 {
				mMat1[i][j] = mat1[i][j]
			}
		}
	}

	// Step 2: Compress mat2 — column-wise non-zero elements
	mMat2 := make([]map[int]int, n)
	for j := 0; j < n; j++ {
		mMat2[j] = make(map[int]int)
		for i := 0; i < k; i++ {
			if mat2[i][j] != 0 {
				mMat2[j][i] = mat2[i][j]
			}
		}
	}

	// Step 3: Multiply only non-zero pairs
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			sum := 0
			for k1, v1 := range mMat1[i] {
				if v2, ok := mMat2[j][k1]; ok {
					sum += v1 * v2
				}
			}
			result[i][j] = sum
		}
	}

	return result
}
