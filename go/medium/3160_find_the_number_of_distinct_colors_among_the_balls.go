// https://leetcode.com/problems/find-the-number-of-distinct-colors-among-the-balls/
package medium

func QueryResults(limit int, queries [][]int) []int {
	result := make([]int, 0)

	mColors := map[int]int{}
	mBalls := map[int]int{}

	for _, q := range queries {
		ball := q[0]
		color := q[1]
		if val, ok := mBalls[ball]; ok {
			prevColor := val
			mColors[prevColor]--

			if mColors[prevColor] == 0 {
				delete(mColors, prevColor)
			}
		}
		mBalls[ball] = color
		mColors[color]++

		result = append(result, len(mColors))
	}

	return result
}
