package java.easy;

import java.util.HashSet;
import java.util.Set;

class HappyNumber {
    public boolean isHappy(int n) {
        if (n == 1 || n == -1) {
            return true;
        }

        Set<Integer> visited = new HashSet<>();
        while (!visited.contains(n)) {
            visited.add(n);
            n = sumOfSquare(n);
            if (n == 1) {
                return true;
            }
        }
        return false;
    }

    private int sumOfSquare(int n) {
        int sum = 0;
        while (n > 0) {
            int mod = n % 10;
            sum += mod*mod;
            n = n/10;
        }

        return sum;
    }
}
