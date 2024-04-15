// https://leetcode.com/problems/game-of-life/

package java.medium;

class GameOfLife {

    // Original | New | State
    // 0 | 0 | 0
    // 1 | 0 | 1
    // 0 | 1 | 2
    // 1 | 1 | 3

    public void gameOfLife(int[][] board) {
        int rows = board.length;
        int cols = board[0].length;

        for (int r = 0; r < rows; r++) {
            for (int c = 0; c < cols; c++) {
                int nei = 0;
                for (int i = r - 1; r < r + 2; i++) {
                    for (int j = c - 1; j < c + 2; j++) {
                        if (i < 0 || j < 0 || (i == r && j == c) || i >= rows || j >= cols) {
                            continue;
                        }
                        if (board[i][j] == 1 || board[i][j] == 3) {
                            nei++;
                        }
                    }

                    if (board[r][c] == 1) { // live cell
                        if (nei == 2 || nei == 3) {
                            board[r][c] = 3;
                        }
                    } else {
                        if (nei == 3) {
                            board[r][c] = 2;
                        }
                    }
                }
            }
        }

        for (int i = 0; i < rows; i++) {
            for (int j = 0; j < cols; j++) {
                if (board[i][j] == 1) {
                    board[i][j] = 0;
                } else if (board[i][j] == 2 || board[i][j] == 3) {
                    board[i][j] = 1;
                }
            }
        }
    }
}