// https://leetcode.com/problems/accounts-merge/
package medium

import "sort"

func accountsMerge(accounts [][]string) [][]string {
	adj := map[string][]string{}
	for _, account := range accounts {
		firstEmail := account[1]
		for _, email := range account[1:] {
			adj[email] = append(adj[email], firstEmail)
			adj[firstEmail] = append(adj[firstEmail], email)
		}
	}

	var result [][]string
	var merged []string
	visited := make(map[string]bool)

	var dfs func(email string)
	dfs = func(email string) {
		visited[email] = true
		merged = append(merged, email)
		for _, nei := range adj[email] {
			if !visited[nei] {
				dfs(nei)
			}
		}
	}

	for _, account := range accounts {
		name := account[0]
		firstEmail := account[1]
		if !visited[firstEmail] {
			merged = []string{name}
			dfs(firstEmail)
			sort.Strings(merged[1:])
			result = append(result, merged)
		}
	}

	return result
}
