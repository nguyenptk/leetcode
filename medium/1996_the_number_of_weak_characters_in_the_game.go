// https://leetcode.com/problems/the-number-of-weak-characters-in-the-game/
package medium

import "sort"

func NumberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] == properties[j][0] {
			return properties[i][1] < properties[j][1]
		}
		return properties[i][0] > properties[j][0]
	})

	result := 0
	defenseMax := 0
	for _, v := range properties {
		if v[1] < defenseMax {
			result++
		} else {
			defenseMax = v[1]
		}

	}

	return result
}
