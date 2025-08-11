// https://leetcode.com/problems/trapping-rain-water
package hard

func Trap(height []int) int {
	result := 0

	l, r := 0, len(height)-1
	lMax := 0
	rMax := 0

	for l < r {
		if height[l] < height[r] {
			lMax = max(lMax, height[l])
			result += lMax - height[l]
			l++
		} else {
			rMax = max(rMax, height[r])
			result += rMax - height[r]
			r--
		}
	}

	return result
}

// Time complexity: O(n)
// Space complexity: O(1)
