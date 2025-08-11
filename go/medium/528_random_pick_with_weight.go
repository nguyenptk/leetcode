// https://leetcode.com/problems/random-pick-with-weight/
package medium

import "math/rand"

type Solution struct {
	prefixSum []int
}

func ConstructorSolution(w []int) Solution {
	prefixSum := make([]int, len(w))
	prefixSum[0] = w[0]
	for i := 1; i < len(w); i++ {
		prefixSum[i] = prefixSum[i-1] + w[i]
	}
	return Solution{
		prefixSum: prefixSum,
	}
}

func (this *Solution) PickIndex() int {
	target := rand.Intn(this.prefixSum[len(this.prefixSum)-1]) + 1
	l, r := 0, len(this.prefixSum)
	for l < r {
		m := l + (r-l)/2
		if target > this.prefixSum[m] {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
