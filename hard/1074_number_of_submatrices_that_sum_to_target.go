// https://leetcode.com/problems/number-of-submatrices-that-currSum-to-target/
package hard

func NumSubmatrixSumTarget(matrix [][]int, target int) int {
	row := len(matrix[0])
	col := len(matrix)
	result := 0
	mp := map[int]int{}

	// The sum of a subarray between i and j
	// replace j by j-1
	for _, v := range matrix {
		for j := 1; j < row; j++ {
			v[j] += v[j-1]
		}
	}

	// Iterate row
	for j := 0; j < row; j++ {
		for k := j; k < row; k++ {
			mp = map[int]int{} // clear map
			mp[0] = 1          // re-initialize map[0] = 1
			currSum := 0       // initialize current sum
			// Iterate col
			for i := 0; i < col; i++ {
				// Check first row
				if j > 0 {
					currSum += matrix[i][k] - matrix[i][j-1]
				} else {
					currSum += matrix[i][k]
				}
				// increase the result by (currSum - target)
				result += mp[currSum-target]
				// increase map of currSum by 1
				mp[currSum]++
			}
		}
	}

	return result
}
