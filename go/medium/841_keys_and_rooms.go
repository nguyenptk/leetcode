// https://leetcode.com/problems/keys-and-rooms
package medium

func CanVisitAllRooms(rooms [][]int) bool {
	seen := make([]bool, len(rooms))
	seen[0] = true
	stack := []int{}
	stack = append(stack, 0)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, n := range rooms[node] {
			if !seen[n] {
				seen[n] = true
				stack = append(stack, n)
			}
		}
	}

	for _, v := range seen {
		if !v {
			return false
		}
	}
	return true
}
