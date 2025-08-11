// https://leetcode.com/problems/minimum-domino-rotations-for-equal-row/
package medium

func MinDominoRotations(tops []int, bottoms []int) int {
	n := len(tops)

	var check func(x int) int
	check = func(x int) int {
		rotationsA, rotationsB := 0, 0
		for i := range n {
			if tops[i] != x && bottoms[i] != x {
				return -1
			} else if tops[i] != x {
				rotationsA++
			} else if bottoms[i] != x {
				rotationsB++
			}
		}
		return min(rotationsA, rotationsB)
	}

	rotations := check(tops[0])
	if rotations != -1 || tops[0] == bottoms[0] {
		return rotations
	}
	return check(bottoms[0])
}
