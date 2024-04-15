// https://leetcode.com/problems/container-with-most-water/
package medium

func MaxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	maxArea := 0
	i := 0
	j := len(height) - 1
	for i < len(height)-1 && j > 0 {
		if height[i] > height[j] {
			maxArea = calArea(height[j], j-i, maxArea)
			j--
		} else {
			maxArea = calArea(height[i], j-i, maxArea)
			i++
		}
	}
	return maxArea
}

func calArea(height, distance, maxArea int) int {
	if maxArea < height*distance {
		return height * distance
	}
	return maxArea
}
