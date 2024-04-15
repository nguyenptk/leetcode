// https://leetcode.com/problems/counting-bits/description/
package java.easy;

class CountingBits {
    public int[] countBits(int n) {
        int[] result = new int[n+1];
        result[0] = 0;

        for (int i = 1; i <= n; i++) {
            int sum = 0;
            int tmp = i;
            while (tmp > 0) {
                sum++;
                tmp &= (tmp-1);
            }
            result[i] = sum;
        }

        return result;
    }
}
