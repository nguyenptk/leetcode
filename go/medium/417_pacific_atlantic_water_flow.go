// https://leetcode.com/problems/pacific-atlantic-water-flow
package medium

func PacificAtlantic(heights [][]int) [][]int {
	rows := len(heights)
	cols := len(heights[0])
	pac := make(map[int]bool)
	atl := make(map[int]bool)

	for c := 0; c < cols; c++ {
		dfsPacificAtlantic(heights, 0, c, pac, heights[0][c])
		dfsPacificAtlantic(heights, rows-1, c, atl, heights[rows-1][c])
	}

	for r := 0; r < rows; r++ {
		dfsPacificAtlantic(heights, r, 0, pac, heights[r][0])
		dfsPacificAtlantic(heights, r, cols-1, atl, heights[r][cols-1])
	}

	result := make([][]int, 0)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if pac[r*cols+c] && atl[r*cols+c] {
				result = append(result, []int{r, c})
			}
		}
	}

	return result
}

func dfsPacificAtlantic(heights [][]int, r, c int, visit map[int]bool, prevHeight int) {
	if visit[r*len(heights[0])+c] ||
		r < 0 ||
		c < 0 ||
		r == len(heights) ||
		c == len(heights[0]) ||
		heights[r][c] < prevHeight {
		return
	}
	visit[r*len(heights[0])+c] = true
	dfsPacificAtlantic(heights, r+1, c, visit, heights[r][c])
	dfsPacificAtlantic(heights, r-1, c, visit, heights[r][c])
	dfsPacificAtlantic(heights, r, c+1, visit, heights[r][c])
	dfsPacificAtlantic(heights, r, c-1, visit, heights[r][c])
}
