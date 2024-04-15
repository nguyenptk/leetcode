package java.medium;

class CoinChangeII {
    public int change(int amount, int[] coins) {
        // Track the ways
        int[] dp = new int[amount+1];
        dp[0] = 1;

        for (int coin : coins) {
            for (int i = 1; i <= amount; i++) {
                if (coin <= i) {
                    dp[i] += dp[i-coin];
                }
            }
        }

        return dp[amount];
    }
}
