// https://leetcode.com/problems/find-missing-and-repeated-values/
package easy

func FindMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	total := n * n

	sumVal := 0
	sqrSum := int64(0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sumVal += grid[i][j]
			sqrSum += int64(grid[i][j] * grid[i][j])
		}
	}

	expectedSum := total * (total + 1) / 2
	expectedSqrSum := int64(total * (total + 1) * (2*total + 1) / 6)

	sumDiff := sumVal - expectedSum
	sqrDiff := sqrSum - expectedSqrSum

	diff := int64(sumDiff)
	sumPlus := sqrDiff / diff

	repeated := int((diff + sumPlus) / 2)
	missing := int((sumPlus - diff) / 2)

	return []int{repeated, missing}
}
