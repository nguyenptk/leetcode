package java.medium;

class SetMatrixZeroes {
    public void setZeroes(int[][] matrix) {
        int row = matrix.length;
        int col = matrix[0].length;
        boolean rowZero = false;

        for (int i = 0; i < row; i++) {
            for (int j = 0; j < col; j++) {
                if (matrix[i][j] == 0) {
                    matrix[0][j] = 0;
                    if (i == 0) {
                        rowZero = true;
                    } else {
                        matrix[i][0] = 0;
                    }
                }
            }
        }

        for (int i = 1; i < row; i++) {
            for (int j = 1; j < col; j++) {
                if (matrix[0][j] == 0 || matrix[i][0] == 0) {
                    matrix[i][j] = 0;
                }
            }
        }

        if (matrix[0][0] == 0) {
            for (int i = 0; i < row; i++) {
                matrix[i][0] = 0;
            }
        }

        if (rowZero) {
            for (int j = 0; j < col; j++) {
                matrix[0][j] = 0;
            }
        }
    }
}
