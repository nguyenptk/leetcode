// https://leetcode.com/problems/flood-fill/
package easy

func FloodFill(image [][]int, sr int, sc int, color int) [][]int {
	oriColor := image[sr][sc]

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= len(image) || j < 0 || j >= len(image[0]) {
			return
		}
		if image[i][j] != oriColor {
			return
		}

		image[i][j] = color
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	if oriColor != color {
		dfs(sr, sc)
	}

	return image
}
