// https://leetcode.com/problems/largest-rectangle-in-histogram
package hard

type Rectangle struct {
	Index  int
	Height int
}

func LargestRectangleArea(heights []int) int {
	maxArea := 0
	stack := make([]Rectangle, len(heights))

	for i, h := range heights {
		start := i
		for len(stack) > 0 && stack[len(stack)-1].Height > h {
			rec := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // POP
			maxArea = max(maxArea, rec.Height*(i-rec.Index))
			start = rec.Index
		}
		stack = append(stack, Rectangle{start, h})
	}

	for len(stack) > 0 {
		rec := stack[len(stack)-1]
		stack = stack[:len(stack)-1] // POP
		maxArea = max(maxArea, rec.Height*(len(heights)-rec.Index))
	}

	return maxArea
}

// func LargestRectangleArea(heights []int) int {
// 	stack := []int{}
// 	stack = append(stack, -1)
// 	maxArea := 0

// 	for i := 0; i < len(heights); i++ {
// 		for stack[len(stack)-1] != -1 && heights[stack[len(stack)-1]] >= heights[i] {
// 			top := stack[len(stack)-1]
// 			stack = stack[:len(stack)-1] // POP
// 			currHeight := heights[top]
// 			currWidth := i - stack[len(stack)-1] - 1
// 			maxArea = max(maxArea, currHeight*currWidth)
// 		}
// 		stack = append(stack, i)
// 	}

// 	for stack[len(stack)-1] != -1 {
// 		top := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1] // POP
// 		currHeight := heights[top]
// 		currWidth := len(heights) - stack[len(stack)-1] - 1
// 		maxArea = max(maxArea, currHeight*currWidth)
// 	}

// 	return maxArea
// }
