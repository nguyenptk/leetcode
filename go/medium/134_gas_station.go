package medium

func CanCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	start := n - 1
	end := 0
	tank := gas[start] - cost[start]

	for start > end {
		if tank < 0 {
			start--
			tank += gas[start] - cost[start]
		} else {
			tank += gas[end] - cost[end]
			end++
		}
	}

	if tank >= 0 {
		return start
	}
	return -1
}
