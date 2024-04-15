// https://leetcode.com/problems/jump-game
package medium

func CanJump(nums []int) bool {
	reach := len(nums) - 1

	for i := reach; i >= 0; i-- {
		if i+nums[i] >= reach {
			reach = i
		}
	}

	return reach == 0
}

type Index int

const (
	GOOD    Index = iota
	BAD     Index = iota
	UNKNOWN Index = iota
)

func CanJumpDPBottomUp(nums []int) bool {
	memo := make([]Index, len(nums))
	for i := range memo {
		memo[i] = UNKNOWN
	}

	memo[len(nums)-1] = GOOD

	for i := len(nums) - 2; i >= 0; i-- {
		// Jump to the second last one
		furthestJump := min(i+nums[i], len(nums)-1)
		for j := i + 1; j <= furthestJump; j++ {
			if memo[j] == GOOD {
				memo[i] = GOOD
				break
			}
		}
	}

	return memo[0] == GOOD
}
