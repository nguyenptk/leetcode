// https://leetcode.com/problems/letter-tile-possibilities/
package medium

func NumTilePossibilities(tiles string) int {
	count := map[byte]int{}
	for i := 0; i < len(tiles); i++ {
		count[tiles[i]]++
	}

	var backtrack func() int
	backtrack = func() int {
		result := 0

		for k, v := range count {
			if v > 0 {
				count[k]--
				result++
				result += backtrack() + 1
				count[k]++
			}
		}

		return result
	}

	return backtrack()
}
