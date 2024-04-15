package java.medium;

class RotateImage {
    public void rotate(int[][] matrix) {
        // int n = matrix.length;
        // for (int i = 0; i < (n+1)/2; i++) {
        //     for (int j = 0; j < n/2; j++) {
        //         int tmp = matrix[n-1-j][i];
        //         matrix[n-1-j][i] = matrix[n-1-i][n-1-j];
        //         matrix[n-1-i][n-1-j] = matrix[j][n-1-i];
        //         matrix[j][n-1-i] = matrix[i][j];
        //         matrix[i][j] = tmp;
        //     }
        // }


        int l = 0;
        int r = matrix.length - 1;

        while (l < r) {
            for (int i = 0; i < r-l; i++) {
                int top = l;
                int bottom = r;

                // Save the top left
                int topLeft = matrix[top][l+i];

                // Move the bottom left to top left
                matrix[top][l+i] = matrix[bottom-i][l];

                // Move the bottom right to bottom left
                matrix[bottom-i][l] = matrix[bottom][r-i];

                // Move the top right to bottom right
                matrix[bottom][r-i] = matrix[top+i][r];

                // Move the top left to top right
                matrix[top+i][r] = topLeft;
            }
            r--;
            l++;
        }
    }
}
