// https://leetcode.com/problems/kth-smallest-amount-with-single-denomination-combination/
package hard

import (
	"sort"
)

func FindKthSmallest(coins []int, k int) int64 {
	sort.Ints(coins)

	if coins[0] == 1 {
		return int64(k)
	}

	// Step 2: Prune redundant coins
	newCoins := []int{}
	for _, coin := range coins {
		redundant := false
		for _, c := range newCoins {
			if coin%c == 0 {
				redundant = true
				break
			}
		}
		if !redundant {
			newCoins = append(newCoins, coin)
		}
	}

	if len(newCoins) == 1 {
		return int64(k) * int64(newCoins[0])
	}

	// Step 4 & 5: DFS LCM generation with signs
	dic := map[int64]int{}
	var dfs func(i int, x int64, sign int)
	dfs = func(i int, x int64, sign int) {
		for j := i + 1; j < len(newCoins); j++ {
			y := lcm(x, int64(newCoins[j]))
			dic[y] += sign
			dfs(j, y, -sign)
		}
	}
	dfs(-1, 1, 1)

	// Step 6: Remove zero entries
	for key, val := range dic {
		if val == 0 {
			delete(dic, key)
		}
	}

	// Step 7: Binary search
	lo := int64(k)
	hi := int64(newCoins[0]) * int64(k)
	for lo <= hi {
		mid := (lo + hi) / 2
		count := int64(0)
		for lcm, sign := range dic {
			count += (mid / lcm) * int64(sign)
		}
		if count >= int64(k) {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// GCD and LCM in int64
func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int64) int64 {
	return a / gcd(a, b) * b
}
