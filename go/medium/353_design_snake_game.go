// https://leetcode.com/problems/design-snake-game/
package medium

type Position struct {
	x, y int
}

type SnakeGame struct {
	width, height int
	snake         []Position        // dequeue
	occupied      map[Position]bool // O(1) collision check
	food          []Position
	foodIdx       int
	dirMap        map[string]Position
}

func ConstructorSnakeGame(width int, height int, food [][]int) SnakeGame {
	start := Position{0, 0}

	foodX := make([]Position, 0)
	for _, f := range food {
		foodX = append(foodX, Position{f[0], f[1]})
	}

	return SnakeGame{
		width:    width,
		height:   height,
		snake:    []Position{start},
		occupied: map[Position]bool{start: true},
		food:     foodX,
		dirMap: map[string]Position{
			"U": {-1, 0},
			"D": {1, 0},
			"L": {0, -1},
			"R": {0, 1},
		},
	}
}

func (f *SnakeGame) Move(direction string) int {
	head := f.snake[0]
	move := f.dirMap[direction]
	newHead := Position{head.x + move.x, head.y + move.y}

	// Remove tail from occupied set unless eating
	tail := f.snake[len(f.snake)-1]
	if f.foodIdx >= len(f.food) || newHead != f.food[f.foodIdx] {
		delete(f.occupied, tail)
		f.snake = f.snake[:len(f.snake)-1] // shrink
	}

	// Check colisions
	if newHead.x < 0 || newHead.x >= f.height ||
		newHead.y < 0 || newHead.y >= f.width {
		return -1 // wall hit
	}
	if f.occupied[newHead] {
		return -1 // self hit
	}

	// Update snake
	f.snake = append([]Position{newHead}, f.snake...)
	f.occupied[newHead] = true

	// Eat food
	if f.foodIdx < len(f.food) && newHead == f.food[f.foodIdx] {
		f.foodIdx++
	}

	return len(f.snake) - 1
}

// Time complexity: O(1)
// Space complexity: O(n)

/**
 * Your SnakeGame object will be instantiated and called as such:
 * obj := Constructor(width, height, food);
 * param_1 := obj.Move(direction);
 */
