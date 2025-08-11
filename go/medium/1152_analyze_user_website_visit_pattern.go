// https://leetcode.com/problems/analyze-user-website-visit-pattern/
package medium

import (
	"sort"
	"strings"
)

type Record struct {
	time    int
	user    string
	website string
}

func MostVisitedPattern(username []string, timestamp []int, website []string) []string {
	// Step 1: Combine and sort records by timestamp
	n := len(username)
	records := make([]Record, n)
	for i := 0; i < n; i++ {
		records[i] = Record{timestamp[i], username[i], website[i]}
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i].time < records[j].time
	})

	// Step 2: Build user → list of websites map
	userWebs := make(map[string][]string)
	for _, r := range records {
		userWebs[r.user] = append(userWebs[r.user], r.website)
	}

	// Step 3: Build all 3-sequence patterns per user (set to avoid duplicates)
	patternUsers := make(map[string]map[string]bool) // pattern → set of users
	for user, webs := range userWebs {
		if len(webs) < 3 {
			continue
		}
		seen := make(map[string]bool)
		for i := 0; i < len(webs)-2; i++ {
			for j := i + 1; j < len(webs)-1; j++ {
				for k := j + 1; k < len(webs); k++ {
					pattern := strings.Join([]string{webs[i], webs[j], webs[k]}, ",")
					if !seen[pattern] {
						if patternUsers[pattern] == nil {
							patternUsers[pattern] = make(map[string]bool)
						}
						patternUsers[pattern][user] = true
						seen[pattern] = true
					}
				}
			}
		}
	}

	// Step 4: Find the pattern with the highest user count
	var result string
	maxCount := 0
	for pattern, users := range patternUsers {
		if len(users) > maxCount || (len(users) == maxCount && pattern < result) {
			maxCount = len(users)
			result = pattern
		}
	}

	return strings.Split(result, ",")
}

// | Operation          | Complexity                                 |
// | ------------------ | ------------------------------------------ |
// | Sort visits        | O(n log n)                                 |
// | Generate sequences | O(u × m³): u = users, m = visits per user  |
// | Space for patterns | O(p × u): p = patterns                     |
