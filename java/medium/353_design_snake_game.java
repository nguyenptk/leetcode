// https://leetcode.com/problems/design-snake-game/

package java.medium;

import java.util.ArrayDeque;
import java.util.Deque;
import java.util.HashSet;
import java.util.Set;

class SnakeGame {

    private int height;
    private int width;
    private int[][] food;
    private int score;
    private int foodIndex;
    private Deque<Integer> snakeBody = new ArrayDeque<>();
    private Set<Integer> visited = new HashSet<>();

    public SnakeGame(int width, int height, int[][] food) {
        this.width = width;
        this.height = height;
        this.food = food;
        snakeBody.offer(0);
        visited.add(0);
    }

    public int move(String direction) {
        int head = snakeBody.peekFirst();
        int row = head / width;
        int col = head % width;
        int newRow = row;
        int newCol = col;

        switch (direction) {
            case "U":
                newRow--;
                break;
            case "D":
                newRow++;
                break;
            case "L":
                newCol--;
                break;
            case "R":
                newCol++;
                break;

            default:
                break;
        }

        // Check if the new head position is out of bounds.
        if (newRow < 0 || newRow >= height || newCol < 0 || newCol >= width) {
            return -1;
        }

        // Check if the snake eats food.
        if (foodIndex < food.length && newRow == food[foodIndex][0] && newCol == food[foodIndex][1]) {
            score++;
            foodIndex++;
        } else {
            // If not eating, move the tail.
            int tail = snakeBody.pollLast();
            visited.remove(tail);
        }

        int newHead = flatternPosition(newRow, newCol);

        // Check if the snake bites itself.
        if (visited.contains(newHead)) {
            return -1;
        }

        // Add the new head to the snake body and mark it visited.
        snakeBody.offerFirst(newHead);
        visited.add(newHead);

        return score;
    }

    private int flatternPosition(int newRow, int newCol) {
        return newRow * width + newCol;
    }
}

/**
 * Your SnakeGame object will be instantiated and called as such:
 * SnakeGame obj = new SnakeGame(width, height, food);
 * int param_1 = obj.move(direction);
 */