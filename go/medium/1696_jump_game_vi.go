// https://leetcode.com/problems/jump-game-vi/
package medium

func MaxResult(nums []int, k int) int {
	n := len(nums)
	a := 0
	b := 0
	deq := make([]int, n)
	deq[0] = n - 1
	for i := n - 2; i >= 0; i-- {
		if deq[a]-i > k {
			a++
		}
		nums[i] += nums[deq[a]]
		for b >= a && nums[deq[b]] <= nums[i] {
			b--
		}
		b++
		deq[b] = i
	}
	return nums[0]
}
