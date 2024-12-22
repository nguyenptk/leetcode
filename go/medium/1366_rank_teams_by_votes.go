// https://leetcode.com/problems/rank-teams-by-votes/

package medium

import "sort"

func RankTeams(votes []string) string {
	ranks := [26][26]int{}
	for _, v := range votes {
		for i, c := range v {
			ranks[c-'A'][i]++
		}
	}

	cand := []byte(votes[0]) // get the list of char of candidates
	sort.Slice(cand, func(i, j int) bool {
		a := ranks[cand[i]-'A']
		b := ranks[cand[j]-'A']
		for r := 0; r < 26; r++ {
			if a[r] == b[r] {
				continue
			}
			return a[r] > b[r]
		}
		return cand[i] < cand[j]
	})

	return string(cand)
}
