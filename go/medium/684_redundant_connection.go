// https://leetcode.com/problems/redundant-connection
package medium

func FindRedundantConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1)

	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}

	for _, e := range edges {
		if find(parent, e[0]) == find(parent, e[1]) {
			return e
		}
		unify(parent, e[0], e[1])
	}

	return []int{}
}

func find(parent []int, num int) int {
	if parent[num] == num {
		return num
	}
	return find(parent, parent[num])
}

func unify(parent []int, x, y int) {
	parent[find(parent, y)] = find(parent, x)
}
