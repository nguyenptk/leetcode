// https://leetcode.com/problems/number-of-1-bits/

package java.easy;

class NumberOfBits {
    public int hammingWeight(int n) {
        int sum = 0;
        while (n > 0) {
            sum++;
            n &= (n-1);
        }
        return sum;
    }
}
