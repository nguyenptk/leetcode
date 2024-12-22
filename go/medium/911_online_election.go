// https://leetcode.com/problems/online-election
package medium

type TopVotedCandidate struct {
	leaders []int
	times   []int
}

func ConstructorTopVotedCandidate(persons []int, times []int) TopVotedCandidate {
	leaders := make([]int, len(persons))
	votes := make(map[int]int)
	curMax := 0
	curLeader := 0

	for i, p := range persons {
		votes[p]++
		if votes[p] >= curMax {
			curMax = votes[p]
			curLeader = p
		}
		leaders[i] = curLeader
	}

	return TopVotedCandidate{
		leaders: leaders,
		times:   times,
	}
}

func (f *TopVotedCandidate) Q(t int) int {
	index := biSearch(f.times, t) - 1
	return f.leaders[index]
}

func biSearch(times []int, t int) int {
	l := 1
	r := len(times)

	for l < r {
		m := l + (r-l)/2
		if times[m] <= t {
			l = m + 1
		} else {
			r = m
		}
	}

	return l
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */
