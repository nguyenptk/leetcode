package java.easy;

class FindWinnerOnATicTacToeGame {
    public String tictactoe(int[][] moves) {
        int n = 3;
        int[] rows = new int[n];
        int[] cols = new int[n];

        int diag = 0;
        int antiDiag = 0;
        int player = 1;

        for (int[] move : moves) {
            int row = move[0];
            int col = move[1];

            rows[row] += player;
            cols[col] += player;

            if (row == col) {
                diag += player;
            }

            if (row + col == n-1) {
                antiDiag += player;
            }

            if (Math.abs(rows[row]) == n || Math.abs(cols[col]) == n ||
                Math.abs(diag) == n || Math.abs(antiDiag) == n) {
                    return player == 1 ? "A" : "B";
            }

            player *= -1;
        }

        if (moves.length + moves[0].length == 9) {
            return "DRAW";
        } else {
            return "PENDING";
        }
    }
}
