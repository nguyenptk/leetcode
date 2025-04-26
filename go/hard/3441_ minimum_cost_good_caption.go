// https://leetcode.com/problems/minimum-cost-good-caption/
package hard

import "math"

func MinCostGoodCaption(s string) string {
	n := len(s)
	if n < 3 {
		return ""
	}

	// dp[i][prev][cnt] = min cost to fix s[i:] given
	//   prev is the previous character index (0..25, or 26 for “none”)
	//   cnt is how many of that prev we’ve placed (capped at 3)
	const NonePrev = 26
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, 27)
		for p := 0; p < 27; p++ {
			dp[i][p] = []int{-1, -1, -1, -1}
		}
	}

	// abs helper
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	// Forward DP: returns min cost
	var f func(i, prev, cnt int) int
	f = func(i, prev, cnt int) int {
		if i == n {
			// At end: the last run must be >= 3
			if cnt < 3 {
				return math.MaxInt32 / 2
			}
			return 0
		}
		if v := dp[i][prev][cnt]; v >= 0 {
			return v
		}
		best := math.MaxInt32 / 2

		// If we have a “full” run or it's the very first character,
		// we may choose ANY next character.
		if cnt >= 3 || prev == NonePrev {
			for ch := 0; ch < 26; ch++ {
				cost := abs(int(s[i]) - (int('a') + ch))
				nextCnt := 1
				if ch == prev {
					// if we pick same char again, that run is “full”
					nextCnt = 3
				}
				cand := cost + f(i+1, ch, nextCnt)
				if cand < best {
					best = cand
				}
			}
		} else {
			// We are in a “short” run (<3), so we MUST continue with prev
			cost := abs(int(s[i]) - (int('a') + prev))
			cand := cost + f(i+1, prev, cnt+1)
			best = cand
		}

		dp[i][prev][cnt] = best
		return best
	}

	// Kick off the DP
	_ = f(0, NonePrev, 0)
	if dp[0][NonePrev][0] >= math.MaxInt32/2 {
		return ""
	}

	// Reconstruction pass
	res := make([]byte, 0, n)
	var i, prev, cnt = 0, NonePrev, 0
	for i < n {
		// current optimal cost
		curCost := dp[i][prev][cnt]

		if cnt >= 3 || prev == NonePrev {
			// we can choose any next char
			var chooseCh int
			for ch := 0; ch < 26; ch++ {
				cost := abs(int(s[i]) - (int('a') + ch))
				nextCnt := 1
				if ch == prev {
					nextCnt = 3
				}
				if curCost == cost+f(i+1, ch, nextCnt) {
					// first matching ch in lex order
					chooseCh = ch
					cnt = nextCnt
					break
				}
			}
			res = append(res, byte('a'+chooseCh))
			prev = chooseCh
		} else {
			// must continue same char
			cost := abs(int(s[i]) - (int('a') + prev))
			res = append(res, byte('a'+prev))
			cnt++
			_ = cost // matches curCost since only one choice
		}
		i++
	}

	return string(res)
}
