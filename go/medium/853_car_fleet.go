// https://leetcode.com/problems/car-fleet
package medium

import "sort"

func CarFleet(target int, position []int, speed []int) int {
	pair := make([][2]int, len(position))

	for i := range position {
		pair[i] = [2]int{position[i], speed[i]}
	}

	sort.Slice(pair, func(i, j int) bool {
		return pair[i][0] > pair[j][0] // Reverse Sorted Order
	})

	stack := make([]float64, 0)

	for _, ps := range pair {
		p, s := float64(ps[0]), float64(ps[1])
		stack = append(stack, (float64(target)-p)/s)
		if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack)
}
