// https://leetcode.com/problems/minimum-number-of-days-to-make-m-bouquets/
package medium

func MinDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}

	l := 0
	r := 0
	for _, b := range bloomDay {
		r = max(r, b)
	}
	result := -1

	getNumBouquets := func(mid int) int {
		nums := 0
		count := 0
		for _, b := range bloomDay {
			if b <= mid {
				count++
			} else {
				count = 0
			}
			if count == k {
				nums++
				count = 0
			}
		}
		return nums
	}

	for l <= r {
		mid := (l + r) / 2
		if getNumBouquets(mid) >= m {
			result = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return result
}
