// https://leetcode.com/problems/count-zero-request-servers/
package medium

import "sort"

type Qry struct {
	time, idx int
}

func countServers(n int, logs [][]int, x int, queries []int) []int {
	// 0) convert 1-based IDs → 0-based
	for _, lg := range logs {
		lg[0]--
	}

	// 1) sort logs by timestamp
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][1] < logs[j][1]
	})

	// 2) sort queries (keeping original indices)
	qs := make([]Qry, len(queries))
	for i, t := range queries {
		qs[i] = Qry{t, i}
	}
	sort.Slice(qs, func(i, j int) bool {
		return qs[i].time < qs[j].time
	})

	counts := make([]int, n) // exactly n servers
	unique := 0              // # of servers with at least one request in window
	res := make([]int, len(queries))

	l, r := 0, 0
	for _, q := range qs {
		start := q.time - x

		// expand r to include logs[r].time <= q.time
		for r < len(logs) && logs[r][1] <= q.time {
			sID := logs[r][0]
			if counts[sID] == 0 {
				unique++
			}
			counts[sID]++
			r++
		}

		// shrink l to exclude logs[l].time < start
		for l < r && logs[l][1] < start {
			sID := logs[l][0]
			counts[sID]--
			if counts[sID] == 0 {
				unique--
			}
			l++
		}

		// servers with no requests = total n minus those in window
		res[q.idx] = n - unique
	}

	return res
}

// Time complexity: O((L + Q) log(L + Q))
// Space complexity: O(n + Q)
