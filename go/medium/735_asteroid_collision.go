// https://leetcode.com/problems/asteroid-collision/
package medium

func AsteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)

	for _, ast := range asteroids {
		alive := true
		for alive && ast < 0 && len(stack) > 0 && stack[len(stack)-1] > 0 {
			top := stack[len(stack)-1]
			if top < -ast {
				stack = stack[:len(stack)-1]
				continue
			} else if top == -ast {
				stack = stack[:len(stack)-1]
			}
			alive = false
		}
		if alive {
			stack = append(stack, ast)
		}
	}

	return stack
}
