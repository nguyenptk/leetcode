// https://leetcode.com/problems/trapping-rain-water
package hard

func Trap(height []int) int {
	result := 0

	l := 0
	r := len(height) - 1

	lMax := 0
	rMax := 0
	for l < r {
		if height[l] < height[r] {
			if lMax < height[l] {
				lMax = height[l]
			} else {
				result += lMax - height[l]
			}
			l++
		} else {
			if rMax < height[r] {
				rMax = height[r]
			} else {
				result += rMax - height[r]
			}
			r--
		}
	}

	return result
}
