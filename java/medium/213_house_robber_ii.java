// https://leetcode.com/problems/house-robber

package java.medium;

class HouseRobbers {
    public int rob(int[] nums) {
        int n = nums.length;

        if (n == 0) {
            return 0;
        }

        if (n == 1) {
            return nums[0];
        }

        int rob1 = rob(nums, 0, n-2);
        int rob2 = rob(nums, 1, n-1);

        return Math.max(rob1, rob2);
    }

    public int rob(int[] nums, int start, int end) {
        int n = end - start + 1;

        int[] dp = new int[n];
        dp[0] = nums[start];
        if (n > 1) {
            dp[1] = Math.max(nums[start], nums[start+1]);
        }

        for (int i = 2; i < n; i++) {
            dp[i] = Math.max(dp[i-1], nums[start+i] + dp[i-2]);
        }

        return dp[n-1];
    }
}
