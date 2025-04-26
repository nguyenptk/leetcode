// https://leetcode.com/problems/bus-routes/
package hard

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	busStop := make(map[int][]int)
	for bus, stops := range routes {
		for _, stop := range stops {
			busStop[stop] = append(busStop[stop], bus)
		}
	}

	q := make([][]int, 0)
	visit := make(map[int]bool)

	for _, bus := range busStop[source] {
		q = append(q, []int{bus, 1})
		visit[bus] = true
	}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		currBus := node[0]
		changes := node[1]

		for _, stop := range routes[currBus] {
			if stop == target {
				return changes
			}
			for _, connected := range busStop[stop] {
				if !visit[connected] {
					q = append(q, []int{connected, changes + 1})
					visit[connected] = true
				}
			}
		}
	}

	return -1
}
