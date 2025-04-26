// https://leetcode.com/problems/equal-row-and-column-pairs/
package medium

func EqualPairs(grid [][]int) int {
	const LEN = 200
	result := 0
	m := make(map[[LEN]int]int)

	for i := 0; i < len(grid); i++ {
		var arr [LEN]int
		copy(arr[:], grid[i])
		m[arr]++
	}

	for i := 0; i < len(grid); i++ {
		arr := [LEN]int{}
		for j := 0; j < len(grid); j++ {
			arr[j] = grid[j][i]
		}
		result += m[arr]
	}

	return result
}
